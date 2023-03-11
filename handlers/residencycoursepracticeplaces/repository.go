package residencycoursepracticeplaces

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items models.ResidencyCoursePracticePlaces) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.ResidencyCoursePracticePlaces) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("name = EXCLUDED.name").
		Set("link = EXCLUDED.link").
		Set("residency_course_practice_place_order = EXCLUDED.residency_course_practice_place_order").
		Set("division_id = EXCLUDED.division_id").
		Set("residency_course_practice_place_group_id = EXCLUDED.residency_course_practice_place_group_id").
		Exec(r.ctx)
	return err
}
