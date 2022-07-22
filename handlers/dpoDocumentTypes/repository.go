package dpoDocumentTypes

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.DpoDocumentType) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.DpoDocumentTypes, error) {
	items := make(models.DpoDocumentTypes, 0)
	err := r.db().NewSelect().Model(&items).
		Relation("DocumentType.Documents.DocumentsScans.Scan").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.DpoDocumentType, error) {
	item := models.DpoDocumentType{}
	err := r.db().NewSelect().Model(&item).Where("id = ?", id).
		Relation("DocumentTypes.Documents.DocumentsScans.Scan").
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.DpoDocumentType{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.DpoDocumentType) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.DpoDocumentTypes) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("document_type_id = EXCLUDED.document_type_id").
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.DpoDocumentType)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}
