package newsImages

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) createMany(items models.NewsImages) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.NewsImage)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.NewsImages) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("description = EXCLUDED.description").
		Set("file_info_id = EXCLUDED.file_info_id").
		Set("news_image_order = EXCLUDED.news_image_order").
		Exec(r.ctx)
	return err
}
