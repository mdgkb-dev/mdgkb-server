package vacancies

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
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

func (r *Repository) create(item *models.Vacancy) (err error) {
	_, err = r.db.NewInsert().Model(item).ExcludeColumn("responses_count", "new_responses_count").Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (item models.VacanciesWithCount, err error) {
	item.Vacancies = make(models.Vacancies, 0)
	query := r.db.NewSelect().Model(&item.Vacancies).
		Relation("VacancyResponses").
		Relation("Division").
		Relation("FormPattern").
		Relation("VacancyDuties").
		Relation("VacancyRequirements").
		Relation("ContactInfo").
		Relation("ContactInfo.Emails").
		Relation("ContactInfo.TelephoneNumbers").
		Relation("ContactDoctor.Human").
		Relation("VacancyResponses.FormValue")
	r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) get(id *string) (*models.Vacancy, error) {
	item := models.Vacancy{}
	err := r.db.NewSelect().
		Model(&item).
		Relation("Division").
		Relation("VacancyDuties", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("vacancy_duties.vacancy_duty_order")
		}).
		Relation("VacancyRequirements", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("vacancy_requirements.vacancy_requirement_order")
		}).
		Relation("ContactInfo").
		Relation("ContactInfo.Emails").
		Relation("ContactInfo.TelephoneNumbers").
		Relation("ContactDoctor.Human").
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
		Relation("FormPattern.PersonalDataAgreement").
		Where("vacancies_view.id = ?", *id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) getBySlug(slug *string) (*models.Vacancy, error) {
	item := models.Vacancy{}
	err := r.db.NewSelect().
		Model(&item).
		Relation("Division").
		Relation("VacancyDuties").
		Relation("VacancyRequirements").
		Relation("ContactInfo").
		Relation("ContactInfo.Emails").
		Relation("ContactInfo.TelephoneNumbers").
		Relation("ContactDoctor.Human").
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
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Vacancy{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Vacancy) (err error) {
	_, err = r.db.NewUpdate().Model(item).ExcludeColumn("responses_count", "new_responses_count").Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Vacancies) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").Model(&items).ExcludeColumn("responses_count", "new_responses_count").
		Set("id = EXCLUDED.id").
		Set("min_salary = EXCLUDED.min_salary").
		Set("max_salary = EXCLUDED.max_salary").
		Set("form_pattern_id = EXCLUDED.form_pattern_id").
		Exec(r.ctx)
	return err
}

func (r *Repository) createResponse(item *models.VacancyResponse) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}
