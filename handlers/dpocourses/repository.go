package dpocourses

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

func (r *Repository) getAll() (item models.DpoCoursesWithCount, err error) {
	item.DpoCourses = make(models.DpoCourses, 0)
	query := r.db().NewSelect().
		Model(&item.DpoCourses).
		Relation("DpoCoursesTeachers.Teacher.Employee.Human").
		Relation("DpoCoursesSpecializations.Specialization").
		Relation("DpoCoursesDates").
		Relation("FormPattern.Fields.File").
		Relation("FormPattern.Fields.ValueType").
		Relation("Specialization")
	r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) get() (*models.DpoCourse, error) {
	item := models.DpoCourse{}
	err := r.db().NewSelect().Model(&item).
		Relation("DpoCoursesTeachers.Teacher.Doctor.Employee.Human").
		Relation("DpoCoursesSpecializations.Specialization").
		Relation("DpoCoursesDates").
		Relation("FormPattern.Fields", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("fields.field_order")
		}).
		Relation("FormPattern.Fields.File").
		Relation("FormPattern.Fields.ValueType").
		Relation("FormPattern.DefaultFormStatus").
		Relation("FormPattern.FormStatusGroup").
		// Relation("FormPattern.PersonalDataAgreement").
		Relation("Specialization").
		Where("dpo_courses_view.? = ?", bun.Safe(r.queryFilter.Col), r.queryFilter.Value).Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.DpoCourse) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.DpoCourse{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.DpoCourse) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.DpoCourses) (err error) {
	_, err = r.db().NewInsert().On("CONFLICT (id) DO UPDATE").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("name = EXCLUDED.name").
		Set("cost = EXCLUDED.cost").
		Set("hours = EXCLUDED.hours").
		Exec(r.ctx)
	return err
}
