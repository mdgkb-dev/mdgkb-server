package employees

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	// _ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) DB() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) SetQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Create(item *models.Employee) (err error) {
	_, err = r.DB().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) GetAll() (item models.EmployeesWithCount, err error) {
	item.Employees = make(models.Employees, 0)
	query := r.DB().NewSelect().Model(&item.Employees).
		Relation("Human").
		Relation("Human.PhotoMini").
		Relation("Head").
		Relation("Doctor").
		Relation("EducationalAcademic")

	r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) Get(slug string) (*models.Employee, error) {
	item := models.Employee{}
	err := r.DB().NewSelect().Model(&item).Where("?TableAlias.id = ?", slug).
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
		Relation("Doctor.DoctorsDivisions.Division").
		Relation("Doctor.Position").
		Relation("EducationalAcademic").
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) Delete(id string) (err error) {
	_, err = r.DB().NewDelete().Model(&models.Employee{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) Update(item *models.Employee) (err error) {
	_, err = r.DB().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
