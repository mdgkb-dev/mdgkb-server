package fileInfos

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.FileInfo) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.FileInfo) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("file_infos.id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.FileInfos) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("original_name = EXCLUDED.original_name").
		Set("file_system_path = EXCLUDED.file_system_path").
		Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.FileInfo) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(item).
		Set("original_name = EXCLUDED.original_name").
		Set("file_system_path = EXCLUDED.file_system_path").
		Exec(r.ctx)
	return err
}

//func (r *Repository) deleteMany(idPool []string) (err error) {
//	_, err = r.db.NewDelete().
//		Model((*models.Document)(nil)).
//		Where("id IN (?)", bun.In(idPool)).
//		Exec(r.ctx)
//	return err
//}
