package vacancies

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Vacancy) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.Vacancies, error) {
	items := make(models.Vacancies, 0)
	err := r.db.NewSelect().Model(&items).Where("vacancies.archived = false").Scan(r.ctx)
	return items, err
}

func (r *Repository) getAllWithResponses() (models.Vacancies, error) {
	items := make(models.Vacancies, 0)
	err := r.db.NewSelect().
		Model(&items).
		Relation("VacancyResponses").
		Relation("Division").
		Relation("VacancyResponses.ContactInfo.Emails").
		Relation("VacancyResponses.ContactInfo.TelephoneNumbers").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Vacancy, error) {
	item := models.Vacancy{}
	err := r.db.NewSelect().
		Model(&item).
		Relation("VacancyResponses.ContactInfo.Emails").
		Relation("VacancyResponses.ContactInfo.TelephoneNumbers").
		Where("id = ?", *id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Vacancy{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Vacancy) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) createResponse(item *models.VacancyResponse) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}
