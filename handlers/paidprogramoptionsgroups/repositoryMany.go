package paidprogramoptionsgroups

import (
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items models.PaidProgramOptionsGroups) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.PaidProgramOptionsGroup)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.PaidProgramOptionsGroups) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Set("name = EXCLUDED.name").
		Set("item_order = EXCLUDED.item_order").
		Set("paid_program_id = EXCLUDED.paid_program_id").
		Model(&items).
		Exec(r.ctx)
	return err
}
