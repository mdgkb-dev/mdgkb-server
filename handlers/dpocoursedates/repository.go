package dpocoursedates

import (
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items models.NmoCoursesDates) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.NmoCourseDates)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.NmoCoursesDates) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("dpo_course_start = EXCLUDED.dpo_course_start").
		Set("dpo_course_end = EXCLUDED.dpo_course_end").
		Set("dpo_course_id = EXCLUDED.dpo_course_id").
		Exec(r.ctx)
	return err
}
