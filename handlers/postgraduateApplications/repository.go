package postgraduateApplications

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) getAll() (models.PostgraduateApplications, error) {
	items := make(models.PostgraduateApplications, 0)
	query := r.db.NewSelect().
		Model(&items).
		Relation("PostgraduateCourse").
		Relation("FieldValues.File").
		Relation("FieldValues.Field").
		Relation("User.Human")

	r.queryFilter.Paginator.CreatePagination(query)
	r.queryFilter.Filter.CreateFilter(query)
	r.queryFilter.Sorter.CreateOrder(query)
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.PostgraduateApplication, error) {
	item := models.PostgraduateApplication{}
	err := r.db.NewSelect().Model(&item).
		Relation("PostgraduateCourse").
		Relation("User.Human").
		Relation("FieldValues.File").
		Relation("FieldValues.Field").
		Where("postgraduate_applications_view.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.PostgraduateApplication) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.PostgraduateApplication{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.PostgraduateApplication) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
