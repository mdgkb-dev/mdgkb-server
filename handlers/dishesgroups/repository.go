package dishesgroups

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	// _ "github.com/go-pg/pg/v10/orm"
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

func (r *Repository) create(item *models.DishesGroup) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items models.DishesGroups, err error) {
	query := r.db().NewSelect().Model(&items).
		Relation("DishSamples.DishesGroup").
		Relation("DishSamples.Image")
	r.queryFilter.HandleQuery(query)
	err = query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.DishesGroup, error) {
	item := models.DishesGroup{}
	err := r.db().NewSelect().Model(&item).
		Where("dishes_groups.id = ?", id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.DishesGroup{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.DishesGroup) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
