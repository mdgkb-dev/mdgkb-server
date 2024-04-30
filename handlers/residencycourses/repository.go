package residencycourses

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) GetAll(c context.Context) (item models.ResidencyCoursesWithCount, err error) {
	item.ResidencyCourses = make(models.ResidencyCourses, 0)
	query := r.helper.DB.IDB(c).NewSelect().
		Model(&item.ResidencyCourses).
		Relation("MainTeacher.Human").
		Relation("ResidencyCoursesSpecializations.Specialization").
		Relation("ResidencyApplications.FormValue.User.Human").
		Relation("ResidencyApplications.FormValue.FieldValues.File").
		Relation("ResidencyApplications.FormValue.FieldValues.Field").
		Relation("ResidencyApplications.FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Relation("ResidencyApplications.ResidencyCourse.ResidencyCoursesSpecializations.Specialization").
		Relation("ResidencyApplications.ResidencyApplicationPointsAchievements.PointsAchievement").
		Relation("FormPattern.Fields.File").
		Relation("FormPattern.Fields.ValueType").
		Relation("FormPattern.DefaultFormStatus").
		Relation("FormPattern.FormStatusGroup").
		Relation("FormPattern.Fields.MaskTokens").
		Relation("StartYear").
		Relation("EndYear")

	r.helper.SQL.ExtractFTSP(c).HandleQuery(query)
	item.Count, err = query.ScanAndCount(c)
	return item, err
}

func (r *Repository) Get(c context.Context) (*models.ResidencyCourse, error) {
	item := models.ResidencyCourse{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("MainTeacher.Human.PhotoMini").
		Relation("MainTeacher.Human.Photo").
		Relation("ResidencyCoursesSpecializations.Specialization").
		Relation("FormPattern.Fields", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("fields.field_order")
		}).
		Relation("Annotation").
		Relation("Program").
		Relation("Plan").
		Relation("Schedule").
		Relation("FormPattern.Fields.File").
		Relation("FormPattern.Fields.ValueType").
		Relation("FormPattern.DefaultFormStatus").
		Relation("FormPattern.FormStatusGroup").
		Relation("FormPattern.Fields.MaskTokens").
		Relation("ResidencyCoursePracticePlaceGroups.ResidencyCoursePracticePlaces.Division").
		Relation("StartYear").
		Relation("EndYear").
		// Where("?TableAlias.? = ?", bun.Safe(r..Col), r..Value).
		Scan(c)
	return &item, err
}

func (r *Repository) Create(c context.Context, item *models.ResidencyCourse) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) Delete(c context.Context, id *string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.ResidencyCourse{}).Where("id = ?", *id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.ResidencyCourse) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) UpsertMany(c context.Context, items models.ResidencyCourses) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("CONFLICT (id) DO UPDATE").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("cost = EXCLUDED.cost").
		Set("free_places = EXCLUDED.free_places").
		Set("paid_places = EXCLUDED.paid_places").
		Set("main_teacher_id = EXCLUDED.main_teacher_id").
		Exec(c)
	return err
}
