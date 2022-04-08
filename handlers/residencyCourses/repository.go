package residencyCourses

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

func (r *Repository) getAll() (models.ResidencyCourses, error) {
	items := make(models.ResidencyCourses, 0)
	query := r.db.NewSelect().
		Model(&items).
		Relation("ResidencyCoursesTeachers.Teacher.Doctor.Human").
		Relation("ResidencyCoursesSpecializations.Specialization").
		Relation("ResidencyCoursesDates").
		Relation("FormPattern.Fields.File").
		Relation("FormPattern.Fields.ValueType")
	
	r.queryFilter.Paginator.CreatePagination(query)
	r.queryFilter.Filter.CreateFilter(query)
	r.queryFilter.Sorter.CreateOrder(query)
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get() (*models.ResidencyCourse, error) {
	item := models.ResidencyCourse{}
	err := r.db.NewSelect().Model(&item).
		Relation("ResidencyCoursesTeachers.Teacher.Doctor.Human").
		Relation("ResidencyCoursesSpecializations.Specialization").
		Relation("ResidencyCoursesDates").
		Relation("DocumentType.Documents.DocumentsScans.Scan").
		Relation("ResidencyCoursePlans.Plan").
		Relation("FormPattern.Fields", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("fields.field_order")
		}).
		Relation("FormPattern.Fields.File").
		Relation("FormPattern.Fields.ValueType").
		Where("residency_courses_view.? = ?", bun.Safe(r.queryFilter.Col), r.queryFilter.Value).Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.ResidencyCourse) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.ResidencyCourse{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.ResidencyCourse) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
