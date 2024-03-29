package formstatuses

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) getAll() (models.FormStatuses, error) {
	items := make(models.FormStatuses, 0)
	query := r.db().NewSelect().
		Model(&items).
		Relation("Icon").
		Relation("FormStatusGroup").
		Relation("FormStatusToFormStatuses.ChildFormStatus.Icon")
	r.queryFilter.HandleQuery(query)
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) GetAllByGroupID(id *string) (models.FormStatuses, error) {
	items := make(models.FormStatuses, 0)
	query := r.db().NewSelect().
		Model(&items).
		Relation("Icon").
		Relation("FormStatusToFormStatuses.ChildFormStatus.Icon").
		Where("form_statuses_view.form_status_group_id = ?", *id)
	r.queryFilter.HandleQuery(query)
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.FormStatus, error) {
	item := models.FormStatus{}
	err := r.db().NewSelect().Model(&item).
		Relation("Icon").
		Relation("FormStatusGroup").
		Relation("FormStatusEmails").
		Relation("FormStatusToFormStatuses.ChildFormStatus.Icon").
		Where("form_statuses_view.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) upsert(item *models.FormStatus) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
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
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.FormStatuses) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("name = EXCLUDED.name").
		Set("label = EXCLUDED.label").
		Set("send_email = EXCLUDED.send_email").
		Set("color = EXCLUDED.color").
		Set("mod_action_name = EXCLUDED.mod_action_name").
		Set("user_action_name = EXCLUDED.user_action_name").
		Set("is_editable = EXCLUDED.is_editable").
		Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.FormStatus{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}
