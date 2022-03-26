package candidateApplications

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

func (r *Repository) getAll() (models.CandidateApplications, error) {
	items := make(models.CandidateApplications, 0)
	query := r.db.NewSelect().
		Model(&items).
		Relation("CandidateApplicationSpecializations.Specializations")

	r.queryFilter.Paginator.CreatePagination(query)
	r.queryFilter.Filter.CreateFilter(query)
	r.queryFilter.Sorter.CreateOrder(query)
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.CandidateApplication, error) {
	item := models.CandidateApplication{}
	err := r.db.NewSelect().Model(&item).
		Relation("CandidateApplicationSpecializations.Specializations").
		Where("candidate_applications.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.CandidateApplication) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.CandidateApplication{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.CandidateApplication) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
