package educationdocumenttypes

import (
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.EducationDocumentType) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.EducationDocumentTypes, error) {
	items := make(models.EducationDocumentTypes, 0)
	err := r.db().NewSelect().Model(&items).
		Relation("PageSection.PageSectionDocuments", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("documents.document_order ASC")
		}).
		Relation("PageSection.PageSectionDocuments.DocumentsScans.Scan").
		Order("education_document_type_order").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.EducationDocumentType, error) {
	item := models.EducationDocumentType{}
	err := r.db().NewSelect().Model(&item).Where("education_document_types.id = ?", id).
		Relation("PageSection.PageSectionDocuments", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("documents.document_order")
		}).
		Relation("PageSection.PageSectionDocuments.DocumentsScans.Scan").
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.EducationDocumentType{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.EducationDocumentType) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.EducationDocumentTypes) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("document_type_id = EXCLUDED.document_type_id").
		Set("education_document_type_order = EXCLUDED.education_document_type_order").
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.EducationDocumentType)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}
