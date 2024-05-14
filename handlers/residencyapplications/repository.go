package residencyapplications

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) GetAll(c context.Context) (item models.ResidencyApplicationsWithCount, err error) {
	item.ResidencyApplications = make(models.ResidencyApplications, 0)
	query := r.helper.DB.IDB(c).NewSelect().
		Model(&item.ResidencyApplications).
		Relation("ResidencyCourse.ResidencyCoursesSpecializations.Specialization").
		Relation("FormValue.FieldValues.File").
		Relation("ResidencyApplicationPointsAchievements.PointsAchievement").
		Relation("FormValue.FieldValues.Field").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Relation("FormValue.User.Human")

	r.helper.SQL.ExtractFTSP(c).HandleQuery(query)
	item.Count, err = query.ScanAndCount(c)
	return item, err
}

func (r *Repository) Get(c context.Context, id *string) (*models.ResidencyApplication, error) {
	item := models.ResidencyApplication{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("Diploma").
		Relation("ResidencyCourse.ResidencyCoursesSpecializations.Specialization").
		Relation("ResidencyCourse.FormPattern.Fields.File").
		Relation("ResidencyCourse.FormPattern.Fields.ValueType").
		Relation("FormValue.User.Human.Contact").
		Relation("FormValue.User.Human.Contact.Address").
		Relation("FormValue.Fields", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("fields.field_order")
		}).
		Relation("FormValue.Fields.File").
		Relation("FormValue.FormValueFiles.File").
		Relation("FormValue.Fields.ValueType").
		Relation("FormValue.FieldValues.File").
		Relation("FormValue.FieldValues.Field.ValueType").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Relation("ResidencyApplicationPointsAchievements.FileInfo").
		Relation("ResidencyApplicationPointsAchievements.PointsAchievement", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("points_achievement.points_achievements_order")
		}).
		Where("residency_applications_view.id = ?", *id).Scan(c)
	return &item, err
}

func (r *Repository) EmailExists(c context.Context, email string, courseID string) (bool, error) {
	exists, err := r.helper.DB.IDB(c).NewSelect().Model((*models.ResidencyApplication)(nil)).
		Join("JOIN form_values ON residency_applications_view.form_value_id = form_values.id").
		Join("JOIN users ON users.id = form_values.user_id and lower(users.email) = lower(?)", email).
		Where("residency_applications_view.residency_course_id = ?", courseID).Exists(c)
	return exists, err
}

func (r *Repository) TypeExists(c context.Context, email string, main bool) (bool, error) {
	exists, err := r.helper.DB.IDB(c).NewSelect().Model((*models.ResidencyApplication)(nil)).
		Where("residency_applications_view.main = ? and lower(residency_applications_view.email) = lower(?)", main, email).
		Exists(c)
	return exists, err
}

func (r *Repository) Create(c context.Context, item *models.ResidencyApplication) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).ExcludeColumn("user_id").Exec(c)
	return err
}

func (r *Repository) Delete(c context.Context, id *string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.ResidencyApplication{}).Where("id = ?", *id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.ResidencyApplication) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).ExcludeColumn("user_id").Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) UpsertMany(c context.Context, items models.ResidencyApplications) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("CONFLICT (id) DO UPDATE").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("residency_course_id = EXCLUDED.residency_course_id").
		Set("points_achievements = EXCLUDED.points_achievements").
		Set("points_entrance = EXCLUDED.points_entrance").
		Set("application_num = EXCLUDED.application_num").
		Exec(c)
	return err
}
