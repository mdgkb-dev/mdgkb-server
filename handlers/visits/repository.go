package visits

import (
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) upsertMany(items models.Visits) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("date = EXCLUDED.date").
		Set("visits_application_id = EXCLUDED.visits_application_id").
		Set("entered = EXCLUDED.entered").
		Set("exited = EXCLUDED.exited").
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.Visit)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}
