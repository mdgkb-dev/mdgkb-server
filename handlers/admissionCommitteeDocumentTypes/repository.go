package admissionCommitteeDocumentTypes

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.AdmissionCommitteeDocumentType) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.AdmissionCommitteeDocumentTypes, error) {
	items := make(models.AdmissionCommitteeDocumentTypes, 0)
	err := r.db.NewSelect().Model(&items).
		Relation("DocumentType.Documents", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("documents.document_order ASC")
		}).
		Relation("DocumentType.Documents.DocumentsScans.Scan").
		Order("admission_committee_document_type_order").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.AdmissionCommitteeDocumentType, error) {
	item := models.AdmissionCommitteeDocumentType{}
	err := r.db.NewSelect().Model(&item).Where("admission_committee_document_types.id = ?", id).
		Relation("DocumentType.Documents", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("documents.document_order")
		}).
		Relation("DocumentType.Documents.DocumentsScans.Scan").
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.AdmissionCommitteeDocumentType{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.AdmissionCommitteeDocumentType) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.AdmissionCommitteeDocumentTypes) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("document_type_id = EXCLUDED.document_type_id").
		Set("admission_committee_document_type_order = EXCLUDED.admission_committee_document_type_order").
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.AdmissionCommitteeDocumentType)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}
