package questions

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Question) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll(published bool) (models.Questions, error) {
	items := make(models.Questions, 0)
	query := r.db.NewSelect().Model(&items).Order("question_date DESC").Order("is_new DESC")
	r.queryFilter.Paginator.Cursor.Column = "question_date"
	r.queryFilter.Paginator.CreatePagination(query)

	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.Question, error) {
	item := models.Question{}
	err := r.db.NewSelect().Model(&item).Relation("User.Human").Where("questions.id = ?", id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Question{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Question) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) readAnswers(userID string) (err error) {
	_, err = r.db.NewUpdate().Model(&models.Question{}).
		Set("answer_is_read = true").
		Where("user_id = ?", userID).
		Exec(r.ctx)
	return err
}

func (r *Repository) changeNewStatus(id string, isNew bool) (err error) {
	_, err = r.db.NewUpdate().Model(&models.Question{}).
		Set("is_new = ?", isNew).
		Where("id = ?", id).
		Exec(r.ctx)
	return err
}

func (r *Repository) publish(id string) (err error) {
	_, err = r.db.NewUpdate().Model(&models.Question{}).
		Set("published = NOT published").
		Set("is_new = false ").
		Where("id = ?", id).
		Exec(r.ctx)
	return err
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}
