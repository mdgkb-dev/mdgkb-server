package medicalProfiles

import (
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"

	_ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.MedicalProfile) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.MedicalProfiles, error) {
	items := make(models.MedicalProfiles, 0)
	query := r.db.NewSelect().Model(&items).
		Relation("MedicalProfilesDivisions.Division").
		Relation("MedicalProfilesNews.News.PreviewImage").
		Order("medical_profiles.name")
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.MedicalProfile, error) {
	item := models.MedicalProfile{}
	err := r.db.NewSelect().Model(&item).
		Relation("MedicalProfilesDivisions.Division").
		Relation("MedicalProfilesNews.News.PreviewImage").
		Relation("MedicalProfilesNews.News.NewsViews").
		Where("medical_profiles.id = ?", id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.MedicalProfile{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.MedicalProfile) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
