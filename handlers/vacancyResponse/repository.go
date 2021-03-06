package vacancyResponse

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.VacancyResponse) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.VacancyResponses, error) {
	items := make(models.VacancyResponses, 0)
	err := r.db.NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.VacancyResponse, error) {
	item := models.VacancyResponse{}
	err := r.db.NewSelect().
		Model(&item).
		Relation("Vacancy").
		Relation("User.Human").
		Where("vacancy_responses.id = ?", id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.VacancyResponse{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.VacancyResponse) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
