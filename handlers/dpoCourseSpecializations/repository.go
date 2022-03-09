package dpoCourseSpecializations

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) createMany(items models.DpoCoursesSpecializations) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.DpoCoursesSpecializations)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.DpoCoursesSpecializations) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("dpo_course_id = EXCLUDED.dpo_course_id").
		Set("specialization_id = EXCLUDED.specialization_id").
		Exec(r.ctx)
	return err
}
