package addressinfos

import (
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items models.PostAddresses) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.PostAddress)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.PostAddresses) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Set("address = EXCLUDED.address").
		Set("description = EXCLUDED.description").
		Model(&items).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.AddressInfo) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Set("region = EXCLUDED.region").
		Set("city = EXCLUDED.city").
		Set("street = EXCLUDED.street").
		Set("building = EXCLUDED.building").
		Set("flat = EXCLUDED.flat").
		Set("zip = EXCLUDED.zip").
		Set("cii = EXCLUDED.cii").
		Model(item).
		Exec(r.ctx)
	return err
}
