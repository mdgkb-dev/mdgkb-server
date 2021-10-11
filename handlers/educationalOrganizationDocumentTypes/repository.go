package educationalOrganizationDocumentTypes

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) getAll() (models.EducationalOrganizationDocumentTypes, error) {
	items := make(models.EducationalOrganizationDocumentTypes, 0)
	err := r.db.NewSelect().Model(&items).
		Relation("EducationalOrganizationDocumentTypeDocuments.Document.FileInfo").Scan(r.ctx)
	return items, err
}

func (r *Repository) createMany(items models.EducationalOrganizationDocumentTypes) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.EducationalOrganizationDocumentType)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.EducationalOrganizationDocumentTypes) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("name = EXCLUDED.name").
		Exec(r.ctx)
	return err
}
