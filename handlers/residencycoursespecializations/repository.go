package residencycoursespecializations

import (
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items models.ResidencyCoursesSpecializations) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.ResidencyCourseSpecialization)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.ResidencyCoursesSpecializations) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("residency_course_id = EXCLUDED.residency_course_id").
		Set("specialization_id = EXCLUDED.specialization_id").
		Set("main = EXCLUDED.main").
		Exec(r.ctx)
	return err
}
