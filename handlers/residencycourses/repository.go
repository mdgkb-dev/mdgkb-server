package residencycourses

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

func (r *Repository) getAll() (item models.ResidencyCoursesWithCount, err error) {
	item.ResidencyCourses = make(models.ResidencyCourses, 0)
	query := r.db().NewSelect().
		Model(&item.ResidencyCourses).
		Relation("ResidencyCoursesTeachers.Teacher.Doctor.Human").
		Relation("ResidencyCoursesSpecializations.Specialization").
		Relation("ResidencyApplications.FormValue.User.Human").
		Relation("ResidencyApplications.FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Relation("ResidencyApplications.ResidencyCourse.ResidencyCoursesSpecializations.Specialization").
		Relation("ResidencyApplications.ResidencyApplicationPointsAchievements.PointsAchievement").
		Relation("FormPattern.Fields.File").
		Relation("FormPattern.Fields.ValueType").
		Relation("StartYear").
		Relation("EndYear")
	r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) get() (*models.ResidencyCourse, error) {
	item := models.ResidencyCourse{}
	err := r.db().NewSelect().Model(&item).
		Relation("ResidencyCoursesTeachers.Teacher.Doctor.Human").
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
		Relation("StartYear").
		Relation("EndYear").
		Where("residency_courses_view.? = ?", bun.Safe(r.queryFilter.Col), r.queryFilter.Value).Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.ResidencyCourse) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.ResidencyCourse{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.ResidencyCourse) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.ResidencyCourses) (err error) {
	_, err = r.db().NewInsert().On("CONFLICT (id) DO UPDATE").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("cost = EXCLUDED.cost").
		Set("free_places = EXCLUDED.free_places").
		Set("paid_places = EXCLUDED.paid_places").
		Exec(r.ctx)
	return err
}
