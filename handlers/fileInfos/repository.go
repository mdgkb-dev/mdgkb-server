package fileInfos

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

//func (r *Repository) deleteMany(idPool []string) (err error) {
//	_, err = r.db.NewDelete().
//		Model((*models.Document)(nil)).
//		Where("id IN (?)", bun.In(idPool)).
//		Exec(r.ctx)
//	return err
//}

func (r *Repository) upsertMany(items models.FileInfos) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("original_name = EXCLUDED.original_name").
		Set("file_system_path = EXCLUDED.file_system_path").
		Exec(r.ctx)
	return err
}
