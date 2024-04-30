package postgraduatecourses

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) GetAll(c context.Context) (item models.PostgraduateCoursesWithCount, err error) {
	item.PostgraduateCourses = make(models.PostgraduateCourses, 0)
	query := r.helper.DB.IDB(c).NewSelect().
		Model(&item.PostgraduateCourses).
		Relation("PostgraduateCoursesTeachers.Teacher.Employee.Human").
		Relation("PostgraduateCoursesSpecializations.Specialization").
		Relation("PostgraduateCoursesDates").
		Relation("FormPattern.Fields.File").
		Relation("FormPattern.Fields.ValueType").
		Relation("QuestionsFile")
	item.Count, err = query.ScanAndCount(c)
	return item, err
}

func (r *Repository) Get(c context.Context) (*models.PostgraduateCourse, error) {
	item := models.PostgraduateCourse{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("PostgraduateCoursesTeachers.Teacher.Employee.Human").
		Relation("PostgraduateCoursesSpecializations.Specialization").
		Relation("PostgraduateCoursesDates").
		Relation("QuestionsFile").
		Relation("ProgramFile").
		Relation("Calendar").
		Relation("Annotation").
		Relation("PageSection.PageSectionDocuments.DocumentsScans.Scan").
		Relation("PostgraduateCoursePlans.Plan").
		Relation("FormPattern.DefaultFormStatus").
		Relation("FormPattern.FormStatusGroup").
		Relation("FormPattern.Fields", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("fields.field_order")
		}).
		Relation("FormPattern.Fields.File").
		Relation("FormPattern.Fields.ValueType").
		Relation("FormPattern.Fields.MaskTokens").
		// Where("postgraduate_courses_view.? = ?", bun.Safe(r..Col), r..Value).
		Scan(c)
	return &item, err
}

func (r *Repository) Create(c context.Context, item *models.PostgraduateCourse) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) Delete(c context.Context, id *string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.PostgraduateCourse{}).Where("id = ?", *id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.PostgraduateCourse) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) UpsertMany(c context.Context, items models.PostgraduateCourses) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("CONFLICT (id) DO UPDATE").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("cost = EXCLUDED.cost").
		Set("years = EXCLUDED.years").
		Exec(c)
	return err
}
