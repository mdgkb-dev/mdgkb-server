package postgraduateapplications

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

func (r *Repository) getAll() (item models.PostgraduateApplicationsWithCount, err error) {
	item.PostgraduateApplications = make(models.PostgraduateApplications, 0)
	query := r.db().NewSelect().
		Model(&item.PostgraduateApplications).
		Relation("PostgraduateCourse.PostgraduateCoursesSpecializations.Specialization").
		Relation("FormValue.FieldValues.File").
		Relation("FormValue.FieldValues.Field").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Relation("FormValue.User.Human")

	r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) get(id *string) (*models.PostgraduateApplication, error) {
	item := models.PostgraduateApplication{}
	err := r.db().NewSelect().Model(&item).
		Relation("PostgraduateCourse.PostgraduateCoursesSpecializations.Specialization").
		Relation("PostgraduateCourse.FormPattern.Fields.File").
		Relation("PostgraduateCourse.FormPattern.Fields.ValueType").
		Relation("FormValue.User.Human").
		Relation("FormValue.Fields.File").
		Relation("FormValue.Fields.ValueType").
		Relation("FormValue.FieldValues.File").
		Relation("FormValue.FieldValues.Field.ValueType").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Where("postgraduate_applications_view.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) emailExists(email string, courseID string) (bool, error) {
	exists, err := r.db().NewSelect().Model((*models.PostgraduateApplication)(nil)).
		Join("JOIN form_values ON postgraduate_applications_view.form_value_id = form_values.id").
		Join("JOIN users ON users.id = form_values.user_id and users.email = ?", email).
		Where("postgraduate_applications_view.postgraduate_course_id = ?", courseID).Exists(r.ctx)
	return exists, err
}

func (r *Repository) create(item *models.PostgraduateApplication) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.PostgraduateApplication{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.PostgraduateApplication) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
