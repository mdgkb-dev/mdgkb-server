package doctorsdivisions

import (
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items models.DoctorsDivisions) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.DoctorDivision)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.DoctorsDivisions) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("doctor_id = EXCLUDED.doctor_id").
		Set("division_id = EXCLUDED.division_id").
		Set("show = EXCLUDED.show").
		Exec(r.ctx)
	return err
}
