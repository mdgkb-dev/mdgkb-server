package dailymenuitems

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

func (r *Repository) create(item *models.DailyMenuItem) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items models.DailyMenuItems, err error) {
	query := r.db().NewSelect().Model(&items)
	r.queryFilter.HandleQuery(query)
	err = query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.DailyMenuItem, error) {
	item := models.DailyMenuItem{}
	err := r.db().NewSelect().Model(&item).
		Where("DailyMenuItems.id = ?", id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.DailyMenuItem{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.DailyMenuItem) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.DailyMenuItems) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("name = EXCLUDED.name").
		Set("price = EXCLUDED.price").
		Set("weight = EXCLUDED.weight").
		Set("additional_weight = EXCLUDED.additional_weight").
		Set("caloric = EXCLUDED.caloric").
		Set("item_order = EXCLUDED.item_order").
		Set("quantity = EXCLUDED.quantity").
		Set("daily_menu_id = EXCLUDED.daily_menu_id").
		Set("dish_sample_id = EXCLUDED.dish_sample_id").
		Set("available = EXCLUDED.available").
		Model(&items).
		Exec(r.ctx)
	return err
}
