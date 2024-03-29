package medicalprofiles

import (
	"context"

	"mdgkb/mdgkb-server/models"
	//_ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) Create(c context.Context, item *models.MedicalProfile) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (models.MedicalProfiles, error) {
	items := make(models.MedicalProfiles, 0)
	q := r.helper.DB.IDB(c).NewSelect().Model(&items).
		Relation("MedicalProfilesDivisions.Division").
		Relation("MedicalProfilesNews.News.PreviewImage").
		Order("medical_profiles.name")
	r.helper.SQL.ExtractFTSP(c).HandleQuery(q)
	err := q.Scan(c)
	return items, err
}

func (r *Repository) Get(c context.Context, id string) (*models.MedicalProfile, error) {
	item := models.MedicalProfile{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("MedicalProfilesDivisions.Division").
		Relation("MedicalProfilesNews.News.PreviewImage").
		Relation("MedicalProfilesNews.News.NewsViews").
		Where("medical_profiles.id = ?", id).
		Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.MedicalProfile{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.MedicalProfile) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}
