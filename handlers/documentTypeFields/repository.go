package documentTypeFields

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) createMany(items models.DocumentTypeFields) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.DocumentTypeField)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.DocumentTypeFields) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("name = EXCLUDED.name").
		Set(`document_type_field_order = EXCLUDED."document_type_field_order"`).
		Set(`document_type_id = EXCLUDED."document_type_id"`).
		Set("value_type_id = EXCLUDED.value_type_id").
		Exec(r.ctx)
	return err
}
