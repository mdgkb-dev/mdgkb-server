package educationalOrganizationDocumentTypeDocuments

import (

	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) createMany(items models.EducationalOrganizationDocumentTypeDocuments) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.EducationalOrganizationDocumentTypeDocument)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.EducationalOrganizationDocumentTypeDocuments) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("document_id = EXCLUDED.document_id").
		Set("educational_organization_document_type_id = EXCLUDED.educational_organization_document_type_id").
		Exec(r.ctx)
	return err
}
