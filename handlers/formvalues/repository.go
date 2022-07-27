package formvalues

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) upsert(item *models.FormValue) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
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
	err := r.db().NewSelect().Model(&item).
		Relation("User.Human").
		Relation("Child.Human").
		Relation("Fields.File").
		Relation("Fields.ValueType").
		Relation("FieldValues.File").
		Relation("FieldValues.Field.ValueType").
		Relation("FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Relation("ResidencyApplication.FormValue.User.Human").
		Relation("ResidencyApplication.FormValue.Fields.File").
		Relation("ResidencyApplication.FormValue.FieldValues.File").
		Relation("ResidencyApplication.FormValue.Fields.ValueType").
		Relation("ResidencyApplication.FormValue.FieldValues.Field.ValueType").
		Relation("ResidencyApplication.ResidencyApplicationPointsAchievements.FileInfo").
		Relation("ResidencyApplication.ResidencyApplicationPointsAchievements.PointsAchievement", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("points_achievement.points_achievements_order")
		}).
		Where("form_values.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) upsertMany(items models.FormValues) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("viewed_by_user = EXCLUDED.viewed_by_user").
		Set("approving_date = EXCLUDED.approving_date").
		Exec(r.ctx)
	return err
}
