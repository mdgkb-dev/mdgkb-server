package experiences

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) createMany(items models.Experiences) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.Experience)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Experiences) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("doctor_id = EXCLUDED.doctor_id").
		Set("experience_start = EXCLUDED.experience_start").
		Set("experience_end = EXCLUDED.experience_end").
		Set("place = EXCLUDED.place").
		Set("position = EXCLUDED.position").
		Exec(r.ctx)
	return err
}
