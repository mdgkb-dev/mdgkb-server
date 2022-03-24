package postgraduateCourseSpecializations

import (
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) createMany(items models.PostgraduateCoursesSpecializations) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.PostgraduateCoursesSpecializations)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.PostgraduateCoursesSpecializations) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("postgraduate_course_id = EXCLUDED.postgraduate_course_id").
		Set("specialization_id = EXCLUDED.specialization_id").
		Exec(r.ctx)
	return err
}
