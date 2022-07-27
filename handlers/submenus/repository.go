package submenus

import (
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items models.SubMenus) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.SubMenu)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.SubMenus) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("name = EXCLUDED.name").
		Set("description = EXCLUDED.description").
		Set("color = EXCLUDED.color").
		Set("hide = EXCLUDED.hide").
		Set("link = EXCLUDED.link").
		Set("sub_menu_order = EXCLUDED.sub_menu_order").
		Set("icon_id = EXCLUDED.icon_id").
		Exec(r.ctx)
	return err
}
