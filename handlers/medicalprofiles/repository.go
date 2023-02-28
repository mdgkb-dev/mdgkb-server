package medicalprofiles

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	//_ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) SetQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) create(item *models.MedicalProfile) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.MedicalProfiles, error) {
	items := make(models.MedicalProfiles, 0)
	query := r.db().NewSelect().Model(&items).
		Relation("MedicalProfilesDivisions.Division").
		Relation("MedicalProfilesNews.News.PreviewImage").
		Order("medical_profiles.name")
	r.queryFilter.HandleQuery(query)
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.MedicalProfile, error) {
	item := models.MedicalProfile{}
	err := r.db().NewSelect().Model(&item).
		Relation("MedicalProfilesDivisions.Division").
		Relation("MedicalProfilesNews.News.PreviewImage").
		Relation("MedicalProfilesNews.News.NewsViews").
		Where("medical_profiles.id = ?", id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.MedicalProfile{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.MedicalProfile) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
