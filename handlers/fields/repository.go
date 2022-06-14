package fields

import (
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Field) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Field) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("file_infos.id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Fields) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("name = EXCLUDED.name").
		Set("required = EXCLUDED.required").
		Set("comment = EXCLUDED.comment").
		Set("field_order = EXCLUDED.field_order").
		Set("form_id = EXCLUDED.form_id").
		Set("form_pattern_id = EXCLUDED.form_pattern_id").
		Set("value_type_id = EXCLUDED.value_type_id").
		Set("file_id = EXCLUDED.file_id").
		Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.Field) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(item).
		Set("name = EXCLUDED.name").
		Set("required = EXCLUDED.required").
		Set("field_order = EXCLUDED.field_order").
		Set("form_id = EXCLUDED.form_id").
		Set("form_pattern_id = EXCLUDED.form_pattern_id").
		Set("value_type_id = EXCLUDED.value_type_id").
		Set("file_id = EXCLUDED.file_id").
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.Field)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}
