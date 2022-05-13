package candidateExams

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

func (r *Repository) getAll() (models.CandidateExams, error) {
	items := make(models.CandidateExams, 0)
	query := r.db.NewSelect().
		Model(&items)

	r.queryFilter.HandleQuery(query)
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.CandidateExam, error) {
	item := models.CandidateExam{}
	err := r.db.NewSelect().Model(&item).
		Relation("FormPattern.Fields", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("fields.field_order")
		}).
		Relation("FormPattern.Fields.File").
		Relation("FormPattern.Fields.ValueType").
		Relation("DocumentType.Documents.DocumentsScans").
		Where("candidate_exams.id = '8800afcc-4139-4285-b552-b78e85d7f0dd'", id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.CandidateExam) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.CandidateExam{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.CandidateExam) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
