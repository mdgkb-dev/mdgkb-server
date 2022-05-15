package vacancies

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/models"

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

func (r *Repository) create(item *models.Vacancy) (err error) {
	_, err = r.db.NewInsert().ExcludeColumn("responses_count", "new_responses_count").Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.Vacancies, error) {
	items := make(models.Vacancies, 0)
	query := r.db.NewSelect().Model(&items).
		Relation("VacancyResponses")
	//Relation("Division").
	//Relation("VacancyDuties").
	//Relation("VacancyRequirements").
	//Relation("ContactInfo").
	//Relation("ContactInfo.Emails").
	//Relation("ContactInfo.TelephoneNumbers").
	//Relation("ContactDoctor.Human").
	//Relation("VacancyResponses.User.Human.ContactInfo.Emails").
	//Relation("VacancyResponses.User.Human.ContactInfo.TelephoneNumbers").
	//Relation("VacancyResponses.VacancyResponsesToDocuments.Document.DocumentsScans.Scan")

	r.queryFilter.HandleQuery(query)

	err := query.Scan(r.ctx)
	return items, err
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
		Relation("VacancyResponses.User.Human.ContactInfo.Emails").
		Relation("VacancyResponses.User.Human.ContactInfo.TelephoneNumbers").
		Relation("VacancyResponses.VacancyResponsesToDocuments.Document.DocumentsScans.Scan").
		Relation("VacancyResponses.VacancyResponsesToDocuments.Document.DocumentType").
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
		Relation("VacancyResponses.User.Human.ContactInfo.Emails").
		Relation("VacancyResponses.User.Human.ContactInfo.TelephoneNumbers").
		Relation("VacancyResponses.VacancyResponsesToDocuments.Document.DocumentsScans.Scan").
		Relation("VacancyResponses.VacancyResponsesToDocuments.Document.DocumentType").
		Where("vacancies.slug = ?", *slug).
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

func (r *Repository) createResponse(item *models.VacancyResponse) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}
