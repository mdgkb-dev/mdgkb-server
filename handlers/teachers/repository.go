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
		Relation("DpoCourses").
		Relation("Doctor.Human").
		Relation("Doctor.Division").
		Relation("Doctor.MedicalProfile").
		Relation("Doctor.Regalias").
		Relation("Doctor.Human.Photo").
		Relation("Doctor.Human.PhotoMini")
	r.queryFilter.HandleQuery(query)
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Teacher, error) {
	item := models.Teacher{}
	err := r.db().NewSelect().Model(&item).
		Relation("Doctor.Human.Photo").
		Relation("Doctor.Human.PhotoMini").
		Relation("Doctor.Human").
		Where("teachers_view.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.Teacher) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Teacher{}).Where("id = ?", *id).Exec(r.ctx)
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
		Set("doctor_id = EXCLUDED.doctor_id").
		Set("position = EXCLUDED.position").
		Model(&items).
		Exec(r.ctx)
	return err
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}
