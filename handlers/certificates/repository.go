package certificates

import (
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) getAll() (models.Certificates, error) {
	items := make(models.Certificates, 0)
	err := r.db().NewSelect().
		Model(&items).
		Relation("Scan").
		Where("doctor_id is null").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) createMany(items models.Certificates) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.Certificate)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Certificates) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("doctor_id = EXCLUDED.doctor_id").
		Set("scan_id = EXCLUDED.scan_id").
		Set("description = EXCLUDED.description").
		Exec(r.ctx)
	return err
}
