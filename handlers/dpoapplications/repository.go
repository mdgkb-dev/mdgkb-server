package dpoapplications

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) getAll() (item models.DpoApplicationsWithCount, err error) {
	item.DpoApplications = make(models.DpoApplications, 0)
	query := r.db().NewSelect().
		Model(&item.DpoApplications).
		Relation("DpoCourse").
		Relation("FormValue.FieldValues.File").
		Relation("FormValue.FieldValues.Field").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Relation("FormValue.User.Human")

	r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) get(id *string) (*models.DpoApplication, error) {
	item := models.DpoApplication{}
	err := r.db().NewSelect().Model(&item).
		Relation("DpoCourse.FormPattern.Fields.File").
		Relation("DpoCourse.FormPattern.Fields.ValueType").
		Relation("FormValue.User.Human").
		Relation("FormValue.Fields.File").
		Relation("FormValue.Fields.ValueType").
		Relation("FormValue.FieldValues.File").
		Relation("FormValue.FieldValues.FieldValuesFiles.FileInfo").
		Relation("FormValue.FieldValues.Field.ValueType").
		Relation("FormValue.FormStatus.FormStatusGroup").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Where("dpo_applications_view.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) emailExists(email string, courseID string) (bool, error) {
	exists, err := r.db().NewSelect().Model((*models.DpoApplication)(nil)).
		Join("JOIN form_values ON dpo_applications_view.form_value_id = form_values.id").
		Join("JOIN users ON users.id = form_values.user_id and users.email = ?", email).
		Where("dpo_applications_view.dpo_course_id = ?", courseID).Exists(r.ctx)
	return exists, err
}

func (r *Repository) create(item *models.DpoApplication) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.DpoApplication{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.DpoApplication) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
