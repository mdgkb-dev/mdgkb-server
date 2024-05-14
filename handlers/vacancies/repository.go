package vacancies

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) Create(c context.Context, item *models.Vacancy) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).ExcludeColumn("responses_count", "new_responses_count").Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (item models.VacanciesWithCount, err error) {
	item.Vacancies = make(models.Vacancies, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&item.Vacancies).
		Relation("VacancyResponses").
		Relation("Division").
		Relation("FormPattern").
		Relation("VacancyDuties").
		Relation("VacancyRequirements").
		Relation("Contact.Emails").
		Relation("Contact.Phones").
		Relation("ContactDoctor.Employee.Human").
		Relation("VacancyResponses.FormValue")

	r.helper.SQL.ExtractFTSP(c).HandleQuery(query)

	item.Count, err = query.ScanAndCount(c)
	return item, err
}

func (r *Repository) Get(c context.Context, id *string) (*models.Vacancy, error) {
	item := models.Vacancy{}
	err := r.helper.DB.IDB(c).NewSelect().
		Model(&item).
		Relation("Division").
		Relation("VacancyDuties", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("vacancy_duties.vacancy_duty_order")
		}).
		Relation("VacancyRequirements", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("vacancy_requirements.vacancy_requirement_order")
		}).
		Relation("Contact.Emails").
		Relation("Contact.Phones").
		Relation("ContactDoctor.Employee.Human").
		Relation("VacancyResponses.FormValue.User.Human").
		Relation("VacancyResponses.FormValue.Fields.File").
		Relation("VacancyResponses.FormValue.Fields.ValueType").
		Relation("VacancyResponses.FormValue.FieldValues.File").
		Relation("VacancyResponses.FormValue.FieldValues.Field.ValueType").
		Relation("VacancyResponses.FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Relation("FormPattern.Fields", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("fields.field_order")
		}).
		Relation("FormPattern.Fields.File").
		Relation("FormPattern.DefaultFormStatus").
		Relation("FormPattern.FormStatusGroup").
		Relation("FormPattern.Fields.ValueType").
		// Relation("FormPattern.PersonalDataAgreement").
		Where("vacancies_view.id = ?", *id).
		Scan(c)
	return &item, err
}

func (r *Repository) GetBySlug(c context.Context, slug *string) (*models.Vacancy, error) {
	item := models.Vacancy{}
	err := r.helper.DB.IDB(c).NewSelect().
		Model(&item).
		Relation("Division").
		Relation("VacancyDuties").
		Relation("VacancyRequirements").
		Relation("ContactInfo").
		Relation("ContactInfo.Emails").
		Relation("ContactInfo.TelephoneNumbers").
		Relation("ContactDoctor.Employee.Human").
		// Relation("VacancyResponses.User.Human.ContactInfo.Emails").
		// Relation("VacancyResponses.User.Human.ContactInfo.TelephoneNumbers").
		Relation("VacancyResponses.FormValue.FieldValues.File").
		Relation("VacancyResponses.FormValue.FieldValues.Field").
		Relation("VacancyResponses.FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Relation("VacancyResponses.FormValue.User.Human").
		Relation("FormPattern.DefaultFormStatus").
		Relation("FormPattern.FormStatusGroup").
		Relation("FormPattern.Fields", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("fields.field_order")
		}).
		Relation("FormPattern.Fields.File").
		Relation("FormPattern.Fields.ValueType").
		Where("vacancies_view.slug = ?", *slug).
		Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id *string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.Vacancy{}).Where("id = ?", *id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.Vacancy) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).ExcludeColumn("responses_count", "new_responses_count").Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) UpsertMany(c context.Context, items models.Vacancies) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("conflict (id) do update").Model(&items).ExcludeColumn("responses_count", "new_responses_count").
		Set("id = EXCLUDED.id").
		Set("min_salary = EXCLUDED.min_salary").
		Set("max_salary = EXCLUDED.max_salary").
		Set("form_pattern_id = EXCLUDED.form_pattern_id").
		Exec(c)
	return err
}

func (r *Repository) CreateResponse(c context.Context, item *models.VacancyResponse) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}
