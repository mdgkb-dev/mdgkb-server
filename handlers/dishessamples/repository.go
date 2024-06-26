package dishessamples

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
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) create(item *models.DishSample) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items models.DishSamples, err error) {
	query := r.db().NewSelect().Model(&items)
	err = query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.DishSample, error) {
	item := models.DishSample{}
	err := r.db().NewSelect().Model(&item).
		Where("dishes_samples.id = ?", id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.DishSample{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.DishSample) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) updateAll(items models.DishSamples) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("item_order = EXCLUDED.item_order").
		Set("quantity = EXCLUDED.quantity").
		Set("weight = EXCLUDED.weight").
		Set("additional_weight = EXCLUDED.additional_weight").
		Set("proteins = EXCLUDED.proteins").
		Set("fats = EXCLUDED.fats").
		Set("carbohydrates = EXCLUDED.carbohydrates").
		Set("dietary = EXCLUDED.dietary").
		Set("lean = EXCLUDED.lean").
		Set("composition = EXCLUDED.composition").
		Set("description = EXCLUDED.description").
		Exec(r.ctx)
	return err
}
