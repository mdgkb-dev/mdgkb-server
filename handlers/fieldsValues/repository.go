package fieldsValues

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.FieldValue) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.FieldValue) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("field_values.id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.FieldValues) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("value_string = EXCLUDED.value_string").
		Set("value_number = EXCLUDED.value_number").
		Set("value_date = EXCLUDED.value_date").
		Set("field_id = EXCLUDED.field_id").
		Set("event_application_id = EXCLUDED.event_application_id").
		Set("dpo_application_id = EXCLUDED.dpo_application_id").
		Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.FieldValue) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(item).
		Set("value_string = EXCLUDED.value_string").
		Set("value_number = EXCLUDED.value_number").
		Set("value_date = EXCLUDED.value_date").
		Set("field_id = EXCLUDED.field_id").
		Set("event_application_id = EXCLUDED.event_application_id").
		Set("dpo_application_id = EXCLUDED.dpo_application_id").
		Exec(r.ctx)
	return err
}

//func (r *Repository) deleteMany(idPool []string) (err error) {
//	_, err = r.db.NewDelete().
//		Model((*models.DocumentType)(nil)).
//		Where("id IN (?)", bun.In(idPool)).
//		Exec(r.ctx)
//	return err
//}
