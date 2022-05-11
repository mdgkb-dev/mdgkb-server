package educationYears

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) getAll() (models.EducationYears, error) {
	items := make(models.EducationYears, 0)
	query := r.db.NewSelect().
		Model(&items)
	r.queryFilter.HandleQuery(query)
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get() (*models.EducationYear, error) {
	item := models.EducationYear{}
	err := r.db.NewSelect().Model(&item).
		Where("education_year.? = ?", bun.Safe(r.queryFilter.Col), r.queryFilter.Value).Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.EducationYear) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.EducationYear{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.EducationYear) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
