package documentTypes

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.DocumentType) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll(params models.DocumentsParams) (items models.DocumentTypes, err error) {
	query := r.db.NewSelect().Model(&items).Relation("DocumentTypeFields.ValueType")
	params.CreateJoin(query)
	err = query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.DocumentType, error) {
	item := models.DocumentType{}
	err := r.db.NewSelect().Model(&item).
		Relation("DocumentTypeFields.ValueType").
		Where("document_types.id = ?", *id).
		Scan(context.Background())
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.DocumentType{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.DocumentType) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.DocumentType) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(item).
		Set("name = EXCLUDED.name").
		Set("description = EXCLUDED.description").
		Set("document_type_order = EXCLUDED.document_type_order").
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.DocumentTypes) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("name = EXCLUDED.name").
		Set("description = EXCLUDED.description").
		Set("document_type_order = EXCLUDED.document_type_order").
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.DocumentType)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}
