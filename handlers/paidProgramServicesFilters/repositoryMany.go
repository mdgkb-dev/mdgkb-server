package paidProgramServicesOptions

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) createMany(items models.PaidProgramServicesOptions) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.PaidProgramServiceFilter)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.PaidProgramServicesOptions) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Set("paid_program_filter_id = EXCLUDED.paid_program_filter_id").
		Set("paid_program_service_id = EXCLUDED.paid_program_service_id").
		Model(&items).
		Exec(r.ctx)
	return err
}
