package dailymenus

import (
	"context"
	"mdgkb/mdgkb-server/models"
	"time"

	"github.com/uptrace/bun"
	// _ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) Create(c context.Context, item *models.DailyMenu) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (models.DailyMenus, error) {
	items := make(models.DailyMenus, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&items).
		Relation("DailyMenuItems", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("daily_menu_items.item_order")
		}).
		Relation("DailyMenuItems.DishesGroup").
		Relation("DailyMenuItems.DishSample.DishesGroup").
		Relation("DailyMenuItems.DishSample.Image")

	r.helper.SQL.ExtractFTSP(c).HandleQuery(query)
	err := query.Scan(c)
	return items, err
}

func (r *Repository) Get(c context.Context, id string) (*models.DailyMenu, error) {
	item := models.DailyMenu{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Where("?TableAlias.id = ?", id).
		Scan(c)
	return &item, err
}

func (r *Repository) GetTodayActive(c context.Context) (*models.DailyMenu, error) {
	item := models.DailyMenu{}
	today := time.Now().Format("2006-01-02")
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("DailyMenuItems", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("daily_menu_items.item_order")
		}).
		Relation("DailyMenuItems.DishesGroup").
		Relation("DailyMenuItems.DishSample.DishesGroup").
		Relation("DailyMenuItems.DishSample.Image").
		Where("?TableAlias.item_date = ?", today).
		Where("?TableAlias.active = true").
		Scan(c)

	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.DailyMenu{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.DailyMenu) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) UpdateAll(c context.Context, items models.DailyMenus) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("item_order = EXCLUDED.item_order").
		Exec(c)
	return err
}
