package employees

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	// _ "github.com/go-pg/pg/v10/orm"
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

func (r *Repository) create(item *models.Employee) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (item models.EmployeesWithCount, err error) {
	item.Employees = make(models.Employees, 0)
	query := r.db().NewSelect().Model(&item.Employees).
		Relation("Human").
		Relation("Human.PhotoMini").
		Relation("Head").
		Relation("Doctor")

	r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) get(slug string) (*models.Employee, error) {
	item := models.Employee{}
	err := r.db().NewSelect().Model(&item).Where("?TableAlias.id = ?", slug).
		Relation("Human.Photo").
		Relation("Human.PhotoMini").
		Relation("Regalias").
		Relation("Experiences").
		Relation("Certificates.Scan").
		Relation("Educations").
		Relation("Certifications").
		Relation("Accreditations").
		Relation("TeachingActivities").
		Relation("Head").
		Relation("Doctor").
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Employee{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Employee) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
