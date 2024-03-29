package dailymenus

import (
	"mdgkb/mdgkb-server/models"
	"time"

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

func (r *Repository) create(item *models.DailyMenu) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.DailyMenus, error) {
	items := make(models.DailyMenus, 0)
	query := r.db().NewSelect().Model(&items).
		Relation("DailyMenuItems", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("daily_menu_items.item_order")
		}).
		Relation("DailyMenuItems.DishesGroup").
		Relation("DailyMenuItems.DishSample.DishesGroup").
		Relation("DailyMenuItems.DishSample.Image")
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.DailyMenu, error) {
	item := models.DailyMenu{}
	err := r.db().NewSelect().Model(&item).
		Where("?TableAlias.id = ?", id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) getTodayActive() (*models.DailyMenu, error) {
	item := models.DailyMenu{}
	today := time.Now().Format("2006-01-02")
	err := r.db().NewSelect().Model(&item).
		Relation("DailyMenuItems", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("daily_menu_items.item_order")
		}).
		Relation("DailyMenuItems.DishesGroup").
		Relation("DailyMenuItems.DishSample.DishesGroup").
		Relation("DailyMenuItems.DishSample.Image").
		Where("?TableAlias.item_date = ?", today).
		Where("?TableAlias.active = true").
		Scan(r.ctx)

	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.DailyMenu{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.DailyMenu) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) updateAll(items models.DailyMenus) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("item_order = EXCLUDED.item_order").
		Exec(r.ctx)
	return err
}
