package educationalacademics

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) getAll() (models.EducationalAcademics, error) {
	items := make(models.EducationalAcademics, 0)
	query := r.db().NewSelect().
		Model(&items).
		Relation("Employee.Human.PhotoMini").
		Relation("Employee.Human.Contact.Emails").
		Relation("Employee.Human.Contact.Phones")
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.EducationalAcademic, error) {
	item := models.EducationalAcademic{}
	err := r.db().NewSelect().Model(&item).
		Relation("Doctor.Employee.Human").
		Where("?TableAlias.id = ?", id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.EducationalAcademic) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.EducationalAcademic{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.EducationalAcademic) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) updateAll(items models.EducationalAcademics) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("item_order = EXCLUDED.item_order").
		Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.EducationalAcademic) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Set("item_order = EXCLUDED.item_order").
		Model(item).
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.EducationalAcademic)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}
