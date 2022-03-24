package postgraduateCourseDates

import (
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) createMany(items models.PostgraduateCoursesDates) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.PostgraduateCoursesDates)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.PostgraduateCoursesDates) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("postgraduate_course_start = EXCLUDED.postgraduate_course_start").
		Set("postgraduate_course_end = EXCLUDED.postgraduate_course_end").
		Set("postgraduate_course_id = EXCLUDED.postgraduate_course_id").
		Exec(r.ctx)
	return err
}
