package educationPublicDocumentTypes

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.EducationPublicDocumentType) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.EducationPublicDocumentTypes, error) {
	items := make(models.EducationPublicDocumentTypes, 0)
	err := r.db().NewSelect().Model(&items).
		Relation("PublicDocumentType.DocumentTypes.Documents.DocumentsScans.Scan").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.EducationPublicDocumentType, error) {
	item := models.EducationPublicDocumentType{}
	err := r.db().NewSelect().Model(&item).Where("id = ?", id).
		Relation("PublicDocumentType.DocumentTypes.Documents.DocumentsScans.Scan").
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.EducationPublicDocumentType{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.EducationPublicDocumentType) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.EducationPublicDocumentType) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(item).
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteByPublicDocumentTypeID(id uuid.NullUUID) (err error) {
	_, err = r.db().NewDelete().Model(&models.EducationPublicDocumentType{}).Where("public_document_type_id = ?", id).Exec(r.ctx)
	return err
}
