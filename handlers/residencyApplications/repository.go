package residencyApplications

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/models"

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

func (r *Repository) getAll() (item models.ResidencyApplicationsWithCount, err error) {
	item.ResidencyApplications = make(models.ResidencyApplications, 0)
	query := r.db().NewSelect().
		Model(&item.ResidencyApplications).
		Relation("ResidencyCourse.ResidencyCoursesSpecializations.Specialization").
		Relation("FormValue.FieldValues.File").
		Relation("ResidencyApplicationPointsAchievements.PointsAchievement").
		Relation("FormValue.FieldValues.Field").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Relation("FormValue.User.Human")
	r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) get(id *string) (*models.ResidencyApplication, error) {
	item := models.ResidencyApplication{}
	err := r.db().NewSelect().Model(&item).
		Relation("ResidencyCourse.ResidencyCoursesSpecializations.Specialization").
		Relation("ResidencyCourse.FormPattern.Fields.File").
		Relation("ResidencyCourse.FormPattern.Fields.ValueType").
		Relation("FormValue.User.Human").
		Relation("FormValue.Fields.File").
		Relation("FormValue.Fields.ValueType").
		Relation("FormValue.FieldValues.File").
		Relation("FormValue.FieldValues.Field.ValueType").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Relation("ResidencyApplicationPointsAchievements.FileInfo").
		Relation("ResidencyApplicationPointsAchievements.PointsAchievement", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("points_achievement.points_achievements_order")
		}).
		Where("residency_applications_view.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) emailExists(email string, courseId string) (bool, error) {
	exists, err := r.db().NewSelect().Model((*models.ResidencyApplication)(nil)).
		Join("JOIN form_values ON residency_applications_view.form_value_id = form_values.id").
		Join("JOIN users ON users.id = form_values.user_id and lower(users.email) = lower(?)", email).
		Where("residency_applications_view.residency_course_id = ?", courseId).Exists(r.ctx)
	return exists, err
}

func (r *Repository) create(item *models.ResidencyApplication) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.ResidencyApplication{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.ResidencyApplication) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.ResidencyApplications) (err error) {
	_, err = r.db().NewInsert().On("CONFLICT (id) DO UPDATE").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("residency_course_id = EXCLUDED.residency_course_id").
		Set("points_achievements = EXCLUDED.points_achievements").
		Set("points_entrance = EXCLUDED.points_entrance").
		Set("application_num = EXCLUDED.application_num").
		Exec(r.ctx)
	return err
}
