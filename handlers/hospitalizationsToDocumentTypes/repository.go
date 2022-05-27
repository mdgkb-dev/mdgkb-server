package hospitalizationsToDocumentTypes

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) createMany(items models.HospitalizationsToDocumentTypes) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.HospitalizationToDocumentType)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.HospitalizationsToDocumentTypes) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Exec(r.ctx)
	return err
}
