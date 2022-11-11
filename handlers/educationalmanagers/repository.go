package educationalmanagers

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

func (r *Repository) getAll() (models.EducationalManagers, error) {
	items := make(models.EducationalManagers, 0)
	query := r.db().NewSelect().
		Model(&items).
		Relation("Doctor.Employee.Human.PhotoMini").
		Relation("Doctor.Employee.Human.ContactInfo.Emails").
		Relation("Doctor.Employee.Human.ContactInfo.TelephoneNumbers")

	r.queryFilter.HandleQuery(query)
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.EducationalManager, error) {
	item := models.EducationalManager{}
	err := r.db().NewSelect().Model(&item).
		Relation("Doctor.Employee.Human").
		Where("educational_managers.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.EducationalManager) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.EducationalManager{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.EducationalManager) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
