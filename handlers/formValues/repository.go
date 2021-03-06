package formValues

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) upsert(item *models.FormValue) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(item).
		Set("id = EXCLUDED.id").
		Set("created_at = EXCLUDED.created_at").
		Set("is_new = EXCLUDED.is_new").
		Set("user_id = EXCLUDED.user_id").
		Set("form_status_id = EXCLUDED.form_status_id").
		Exec(r.ctx)
	return err
}

func (r *Repository) get(id *string) (*models.FormValue, error) {
	item := models.FormValue{}
	err := r.db.NewSelect().Model(&item).
		Relation("User.Human").
		Relation("Fields.File").
		Relation("Fields.ValueType").
		Relation("FieldValues.File").
		Relation("FieldValues.Field.ValueType").
		Relation("FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Where("form_values.id = ?", *id).Scan(r.ctx)
	return &item, err
}
