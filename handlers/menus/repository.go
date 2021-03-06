package menus

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Menu) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.Menus, error) {
	items := make(models.Menus, 0)
	err := r.db.NewSelect().Model(&items).
		Relation("Page").
		Relation("Icon").
		Relation("SubMenus", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("sub_menus.sub_menu_order")
		}).
		Relation("SubMenus.Page").
		Relation("SubMenus.Icon").
		Order("menus.menu_order").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Menu, error) {
	item := models.Menu{}
	err := r.db.NewSelect().Model(&item).
		Relation("Page").
		Relation("Icon").
		Relation("SubMenus", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("sub_menus.sub_menu_order")
		}).
		Relation("SubMenus.Page").
		Relation("SubMenus.Icon").
		Where("menus.id = ?", *id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Menu{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Menu) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.Menu)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Menus) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("menu_order = EXCLUDED.menu_order").
		Set("link = EXCLUDED.link").
		Set("top = EXCLUDED.top").
		Set("side = EXCLUDED.side").
		Exec(r.ctx)
	return err
}
