package educationyears

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	return nil
}

func (r *Repository) getAll() (models.EducationYears, error) {
	items := make(models.EducationYears, 0)
	query := r.db().NewSelect().
		Model(&items)
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get() (*models.EducationYear, error) {
	item := models.EducationYear{}
	err := r.db().NewSelect().Model(&item).
		// Where("education_year.? = ?", bun.Safe(r..Col), r..Value).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.EducationYear) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.EducationYear{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.EducationYear) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
