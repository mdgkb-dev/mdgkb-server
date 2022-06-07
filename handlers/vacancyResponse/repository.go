package vacancyResponse

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) GetDB() *bun.DB {
	return r.db
}

func (r *Repository) SetQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) create(item *models.VacancyResponse) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (item models.VacancyResponsesWithCount, err error) {
	item.VacancyResponses = make(models.VacancyResponses, 0)
	query := r.db.NewSelect().Model(&item.VacancyResponses).
		Relation("Vacancy").
		Relation("FormValue.User.Human").
		Relation("FormValue.Fields.ValueType").
		Relation("FormValue.FieldValues.File").
		Relation("FormValue.FieldValues.Field.ValueType").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus")
	r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
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
		Where("vacancy_responses_view.id = ?", id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.VacancyResponse{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.VacancyResponse)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
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
		Where("vacancy_responses_view.vacancy_id = ?", vacancyId).Exists(r.ctx)
	return exists, err
}
