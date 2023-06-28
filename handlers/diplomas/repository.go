package diplomas

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) upsert(item *models.Diploma) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(item).
		Set("id = EXCLUDED.id").
		Set("university_name = EXCLUDED.university_name").
		Set("university_end_date = EXCLUDED.university_end_date").
		Set("number = EXCLUDED.number").
		Set("series = EXCLUDED.series").
		Set("date = EXCLUDED.date").
		Set("speciality_code = EXCLUDED.speciality_code").
		Set("speciality = EXCLUDED.speciality").
		Exec(r.ctx)
	return err
}
