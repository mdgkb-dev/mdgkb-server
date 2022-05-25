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
		Relation("FormValue.User.Human").
		Relation("FormValue.Fields.ValueType").
		Relation("FormValue.FieldValues.File").
		Relation("FormValue.FieldValues.Field.ValueType").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
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

func (r *Repository) emailExists(email string, vacancyId string) (bool, error) {
	exists, err := r.db.NewSelect().Model((*models.VacancyResponse)(nil)).
		Join("JOIN form_values ON vacancy_responses.form_value_id = form_values.id").
		Join("JOIN users ON users.id = form_values.user_id and users.email = ?", email).
		Where("vacancy_responses.vacancy_id = ?", vacancyId).Exists(r.ctx)
	return exists, err
}
