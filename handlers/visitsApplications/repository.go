package visitsApplications

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func (r *Repository) DB() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) SetQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) getAll() (item models.VisitsApplicationsWithCount, err error) {
	item.VisitsApplications = make(models.VisitsApplications, 0)
	query := r.DB().NewSelect().Model(&item.VisitsApplications).
		Relation("Gate").
		Relation("Division").
		Relation("FormValue.Child.Human").
		Relation("FormValue.User.Human").
		Relation("Visits").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus")
	r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) get(id *string) (*models.VisitsApplication, error) {
	item := models.VisitsApplication{}
	err := r.DB().NewSelect().Model(&item).
		Relation("Gate").
		Relation("Division").
		Relation("FormValue.Child.Human").
		Relation("FormValue.User.Human").
		Relation("FormValue.Fields.ValueType").
		Relation("FormValue.FieldValues.File").
		Relation("FormValue.FieldValues.Field.ValueType").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Relation("Visits").
		Where("visits_applications_view.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.VisitsApplication) (err error) {
	_, err = r.DB().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.DB().NewDelete().Model(&models.VisitsApplication{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.VisitsApplication) (err error) {
	_, err = r.DB().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
