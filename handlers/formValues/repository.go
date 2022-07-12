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
		Set("viewed_by_user = EXCLUDED.viewed_by_user").
		Set("user_id = EXCLUDED.user_id").
		Set("form_status_id = EXCLUDED.form_status_id").
		Set("approving_date = EXCLUDED.approving_date").
		Set("child_id = EXCLUDED.child_id").
		Set("mod_comment = EXCLUDED.mod_comment").
		Exec(r.ctx)
	return err
}

func (r *Repository) get(id *string) (*models.FormValue, error) {
	item := models.FormValue{}
	err := r.db.NewSelect().Model(&item).
		Relation("User.Human").
		Relation("Child.Human").
		Relation("Fields.File").
		Relation("Fields.ValueType").
		Relation("FieldValues.File").
		Relation("FieldValues.Field.ValueType").
		Relation("FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Where("form_values.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) upsertMany(items models.FormValues) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("viewed_by_user = EXCLUDED.viewed_by_user").
		Set("approving_date = EXCLUDED.approving_date").
		Exec(r.ctx)
	return err
}
