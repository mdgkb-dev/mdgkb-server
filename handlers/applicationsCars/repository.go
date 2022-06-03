package applicationsCars

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func (r *Repository) GetDB() *bun.DB {
	return r.db
}

func (r *Repository) SetQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) getAll() (item models.ApplicationsCarsWithCount, err error) {
	item.ApplicationsCars = make(models.ApplicationsCars, 0)
	query := r.db.NewSelect().Model(&item.ApplicationsCars).
		Relation("Gate").
		Relation("Division").
		// Relation("FormValue.Child.Human").
		Relation("FormValue.User.Human").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus")
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) get(id *string) (*models.ApplicationCar, error) {
	item := models.ApplicationCar{}
	err := r.db.NewSelect().Model(&item).
		Relation("Gate").
		Relation("Division").
		Relation("FormValue.Child.Human").
		Relation("FormValue.User.Human").
		Relation("FormValue.Fields.ValueType").
		Relation("FormValue.FieldValues.File").
		Relation("FormValue.FieldValues.Field.ValueType").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Where("applications_cars.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.ApplicationCar) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.ApplicationCar{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.ApplicationCar) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
