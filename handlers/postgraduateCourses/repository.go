package postgraduateCourses

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) getAll() (item models.PostgraduateCoursesWithCount, err error) {
	query := r.db.NewSelect().
		Model(&item.PostgraduateCourses).
		Relation("PostgraduateCoursesTeachers.Teacher.Doctor.Human").
		Relation("PostgraduateCoursesSpecializations.Specialization").
		Relation("PostgraduateCoursesDates").
		Relation("FormPattern.Fields.File").
		Relation("FormPattern.Fields.ValueType").
		Relation("QuestionsFile")
	r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) get() (*models.PostgraduateCourse, error) {
	item := models.PostgraduateCourse{}
	err := r.db.NewSelect().Model(&item).
		Relation("PostgraduateCoursesTeachers.Teacher.Doctor.Human").
		Relation("PostgraduateCoursesSpecializations.Specialization").
		Relation("PostgraduateCoursesDates").
		Relation("QuestionsFile").
		Relation("ProgramFile").
		Relation("Calendar").
		Relation("Annotation").
		Relation("DocumentType.Documents.DocumentsScans.Scan").
		Relation("PostgraduateCoursePlans.Plan").
		Relation("FormPattern.Fields", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("fields.field_order")
		}).
		Relation("FormPattern.Fields.File").
		Relation("FormPattern.Fields.ValueType").
		Where("postgraduate_courses_view.? = ?", bun.Safe(r.queryFilter.Col), r.queryFilter.Value).Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.PostgraduateCourse) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.PostgraduateCourse{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.PostgraduateCourse) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.PostgraduateCourses) (err error) {
	_, err = r.db.NewInsert().On("CONFLICT (id) DO UPDATE").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("cost = EXCLUDED.cost").
		Set("years = EXCLUDED.years").
		Exec(r.ctx)
	return err
}
