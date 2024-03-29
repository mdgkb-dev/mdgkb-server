package postgraduatecourseteachers

import (
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items models.PostgraduateCoursesTeachers) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.PostgraduateCoursesTeachers)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.PostgraduateCoursesTeachers) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("teacher_id = EXCLUDED.teacher_id").
		Set("postgraduate_course_id = EXCLUDED.postgraduate_course_id").
		Set("main = EXCLUDED.main").
		Exec(r.ctx)
	return err
}
