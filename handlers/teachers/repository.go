package teachers

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) getAll() (models.Teachers, error) {
	items := make(models.Teachers, 0)
	query := r.db().NewSelect().Model(&items).
		Relation("NmoCourses").
		Relation("Employee.Human").
		Relation("Employee.Regalias").
		Relation("Employee.Human.Photo").
		Relation("Employee.Human.PhotoMini")
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(slug string) (*models.Teacher, error) {
	item := models.Teacher{}
	err := r.db().NewSelect().Model(&item).Where("teachers_view.slug = ?", slug).
		Relation("Employee.Human.Photo").
		Relation("Employee.Human.PhotoMini").
		Relation("Employee.Human").
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.Teacher) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Teacher{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Teacher) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.Teacher)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Teachers) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("employee_id = EXCLUDED.employee_id").
		Set("position = EXCLUDED.position").
		Model(&items).
		Exec(r.ctx)
	return err
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	if err != nil {
		return err
	}
	return nil
}
