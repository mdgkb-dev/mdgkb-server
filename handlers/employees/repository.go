package employees

import (
	"context"
	"mdgkb/mdgkb-server/models"
	// _ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) Create(c context.Context, item *models.Employee) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (item models.EmployeesWithCount, err error) {
	item.Employees = make(models.Employees, 0)
	q := r.helper.DB.IDB(c).NewSelect().Model(&item.Employees).
		Relation("Human").
		Relation("Human.PhotoMini").
		Relation("Head").
		Relation("Doctor").
		Relation("EducationalAcademic")

	r.helper.SQL.ExtractFTSP(c).HandleQuery(q)
	item.Count, err = q.ScanAndCount(c)
	return item, err
}

func (r *Repository) Get(c context.Context, slug string) (*models.Employee, error) {
	item := models.Employee{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).Where("?TableAlias.id = ?", slug).
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
		Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.Employee{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.Employee) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}
