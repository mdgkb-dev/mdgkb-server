package questions

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) create(item *models.Question) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (item models.QuestionsWithCount, err error) {
	item.Questions = make(models.Questions, 0)
	query := r.db().NewSelect().
		Model(&item.Questions).
		Relation("User.Human")

	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) get(id string) (*models.Question, error) {
	item := models.Question{}
	err := r.db().NewSelect().
		Model(&item).
		Relation("User.Human").
		Relation("File").
		Where("questions.id = ?", id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Question{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Question) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) readAnswers(userID string) (err error) {
	_, err = r.db().NewUpdate().Model(&models.Question{}).
		Set("answer_is_read = true").
		Where("user_id = ? and (original_answer= '') IS NOT TRUE", userID).
		Exec(r.ctx)
	return err
}

func (r *Repository) changeNewStatus(id string, isNew bool) (err error) {
	_, err = r.db().NewUpdate().Model(&models.Question{}).
		Set("is_new = ?", isNew).
		Where("id = ?", id).
		Exec(r.ctx)
	return err
}

func (r *Repository) publish(id string) (err error) {
	_, err = r.db().NewUpdate().Model(&models.Question{}).
		Set("published = NOT published").
		Set("is_new = false").
		Where("id = ?", id).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Questions) (err error) {
	_, err = r.db().NewInsert().On("CONFLICT (id) DO UPDATE").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("published = EXCLUDED.published").
		Exec(r.ctx)
	return err
}
