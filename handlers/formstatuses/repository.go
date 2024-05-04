package formstatuses

import (
	"context"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) GetAll(c context.Context) (models.FormStatuses, error) {
	items := make(models.FormStatuses, 0)
	query := r.helper.DB.IDB(c).NewSelect().
		Model(&items).
		Relation("Icon").
		Relation("FormStatusGroup").
		Relation("FormStatusToFormStatuses.ChildFormStatus.Icon")
	err := query.Scan(c)
	return items, err
}

func (r *Repository) GetAllByGroupID(c context.Context, id *string) (models.FormStatuses, error) {
	items := make(models.FormStatuses, 0)
	query := r.helper.DB.IDB(c).NewSelect().
		Model(&items).
		Relation("Icon").
		Relation("FormStatusToFormStatuses.ChildFormStatus.Icon").
		Where("form_statuses_view.form_status_group_id = ?", *id)
	err := query.Scan(c)
	return items, err
}

func (r *Repository) Get(c context.Context, id *string) (*models.FormStatus, error) {
	item := models.FormStatus{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("Icon").
		Relation("FormStatusGroup").
		Relation("FormStatusEmails").
		Relation("FormStatusToFormStatuses.ChildFormStatus.Icon").
		Where("form_statuses_view.id = ?", *id).Scan(c)
	return &item, err
}

func (r *Repository) Upsert(c context.Context, item *models.FormStatus) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("conflict (id) do update").
		Model(item).
		Set("id = EXCLUDED.id").
		Set("name = EXCLUDED.name").
		Set("label = EXCLUDED.label").
		Set("color = EXCLUDED.color").
		Set("send_email = EXCLUDED.send_email").
		Set("mod_action_name = EXCLUDED.mod_action_name").
		Set("user_action_name = EXCLUDED.user_action_name").
		Set("is_editable = EXCLUDED.is_editable").
		Set("form_status_group_id = EXCLUDED.form_status_group_id").
		Set("icon_id = EXCLUDED.icon_id").
		Exec(c)
	return err
}

func (r *Repository) UpsertMany(c context.Context, items models.FormStatuses) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("name = EXCLUDED.name").
		Set("label = EXCLUDED.label").
		Set("send_email = EXCLUDED.send_email").
		Set("color = EXCLUDED.color").
		Set("mod_action_name = EXCLUDED.mod_action_name").
		Set("user_action_name = EXCLUDED.user_action_name").
		Set("is_editable = EXCLUDED.is_editable").
		Exec(c)
	return err
}

func (r *Repository) Delete(c context.Context, id *string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.FormStatus{}).Where("id = ?", *id).Exec(c)
	return err
}
