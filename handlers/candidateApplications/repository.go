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
		Relation("CandidateExam").
		Relation("FormValue.FieldValues.File").
		Relation("FormValue.FieldValues.Field").
		Relation("FormValue.User.Human").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Relation("CandidateApplicationSpecializations.Specialization")

	r.queryFilter.HandleQuery(query)
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.CandidateApplication, error) {
	item := models.CandidateApplication{}
	err := r.db.NewSelect().Model(&item).
		Relation("CandidateApplicationSpecializations.Specialization").
		Relation("CandidateExam.FormPattern.Fields.File").
		Relation("CandidateExam.FormPattern.Fields.ValueType").
		Relation("FormValue.User.Human").
		Relation("FormValue.Fields.File").
		Relation("FormValue.Fields.ValueType").
		Relation("FormValue.FieldValues.File").
		Relation("FormValue.FieldValues.Field.ValueType").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Where("candidate_applications.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) emailExists(email string, examId string) (bool, error) {
	exists, err := r.db.NewSelect().Model((*models.CandidateApplication)(nil)).
		Join("JOIN users ON users.email = ?", email).
		Where("candidate_applications.candidate_exam_id = ?", examId).Exists(r.ctx)
	return exists, err
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
