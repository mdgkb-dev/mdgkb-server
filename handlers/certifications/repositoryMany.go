package certifications

import (
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items models.Certifications) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.Certification)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Certifications) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Set("employee_id = EXCLUDED.employee_id").
		Set("specialization = EXCLUDED.specialization").
		Set("certification_date = EXCLUDED.certification_date").
		Set("end_date = EXCLUDED.end_date").
		Set("document = EXCLUDED.document").
		Set("place = EXCLUDED.place").
		Model(&items).
		Exec(r.ctx)
	return err
}
