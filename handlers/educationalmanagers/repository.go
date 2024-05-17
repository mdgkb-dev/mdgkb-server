package educationalmanagers

import (
	"context"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) GetAll(c context.Context) (models.EducationalManagers, error) {
	items := make(models.EducationalManagers, 0)
	query := r.helper.DB.IDB(c).NewSelect().
		Model(&items).
		Relation("Doctor.Employee.Human.PhotoMini").
		Relation("Doctor.Employee.Human.Contact.Emails").
		Relation("Doctor.Employee.Human.Contact.Phones")

	r.helper.SQL.ExtractFTSP(c).HandleQuery(query)
	err := query.Scan(c)
	return items, err
}

func (r *Repository) Get(c context.Context, id string) (*models.EducationalManager, error) {
	item := models.EducationalManager{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("Doctor.Employee.Human").
		Where("educational_managers.id = ?", id).Scan(c)
	return &item, err
}

func (r *Repository) Create(c context.Context, item *models.EducationalManager) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.EducationalManager{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.EducationalManager) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}
