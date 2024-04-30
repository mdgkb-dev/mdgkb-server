package vacancyresponses

import (
	"context"
	"fmt"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) Create(c context.Context, item *models.VacancyResponse) (err error) {
	fmt.Printf("%+v\n", item)
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (item models.VacancyResponsesWithCount, err error) {
	item.VacancyResponses = make(models.VacancyResponses, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&item.VacancyResponses).
		Relation("Vacancy").
		Relation("FormValue.User.Human").
		Relation("FormValue.Fields.ValueType").
		Relation("FormValue.FieldValues.File").
		Relation("FormValue.FieldValues.Field.ValueType").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus")
	item.Count, err = query.ScanAndCount(c)
	return item, err
}

func (r *Repository) Get(c context.Context, id string) (*models.VacancyResponse, error) {
	item := models.VacancyResponse{}
	err := r.helper.DB.IDB(c).NewSelect().
		Model(&item).
		Relation("Vacancy").
		Relation("FormValue.User.Human.ContactInfo").
		Relation("FormValue.User.Human.ContactInfo.AddressInfo").
		Relation("FormValue.Fields", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("fields.field_order")
		}).
		Relation("FormValue.Fields.File").
		Relation("FormValue.FormValueFiles.File").
		Relation("FormValue.Fields.ValueType").
		Relation("FormValue.FieldValues.File").
		Relation("FormValue.FieldValues.Field.ValueType").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Where("vacancy_responses_view.id = ?", id).Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.VacancyResponse{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) DeleteMany(c context.Context, idPool []uuid.UUID) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().
		Model((*models.VacancyResponse)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.VacancyResponse) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) EmailExists(c context.Context, email string, vacancyID string) (bool, error) {
	exists, err := r.helper.DB.IDB(c).NewSelect().Model((*models.VacancyResponse)(nil)).
		Join("JOIN form_values ON vacancy_responses_view.form_value_id = form_values.id").
		Join("JOIN users ON users.id = form_values.user_id and users.email = ?", email).
		Where("vacancy_responses_view.vacancy_id = ?", vacancyID).Exists(c)
	return exists, err
}
