package forms

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.Form) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Form) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("file_infos.id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Forms) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("id = EXCLUDED.id").
		Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.Form) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(item).
		Set("id = EXCLUDED.id").
		Exec(r.ctx)
	return err
}

//func (r *Repository) deleteMany(idPool []string) (err error) {
//	_, err = r.db().NewDelete().
//		Model((*models.PageSection)(nil)).
//		Where("id IN (?)", bun.In(idPool)).
//		Exec(r.ctx)
//	return err
//}
