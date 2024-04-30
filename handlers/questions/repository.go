package questions

import (
	"context"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) Create(c context.Context, item *models.Question) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (item models.QuestionsWithCount, err error) {
	item.Questions = make(models.Questions, 0)
	query := r.helper.DB.IDB(c).NewSelect().
		Model(&item.Questions).
		Relation("User.Human")

	r.helper.SQL.ExtractFTSP(c).HandleQuery(query)
	item.Count, err = query.ScanAndCount(c)
	return item, err
}

func (r *Repository) Get(c context.Context, id string) (*models.Question, error) {
	item := models.Question{}
	err := r.helper.DB.IDB(c).NewSelect().
		Model(&item).
		Relation("User.Human").
		Relation("File").
		Where("questions.id = ?", id).Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.Question{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.Question) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) ReadAnswers(c context.Context, userID string) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(&models.Question{}).
		Set("answer_is_read = true").
		Where("user_id = ? and (original_answer= '') IS NOT TRUE", userID).
		Exec(c)
	return err
}

func (r *Repository) ChangeNewStatus(c context.Context, id string, isNew bool) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(&models.Question{}).
		Set("is_new = ?", isNew).
		Where("id = ?", id).
		Exec(c)
	return err
}

func (r *Repository) Publish(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(&models.Question{}).
		Set("published = NOT published").
		Set("is_new = false").
		Where("id = ?", id).
		Exec(c)
	return err
}

func (r *Repository) UpsertMany(c context.Context, items models.Questions) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("CONFLICT (id) DO UPDATE").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("published = EXCLUDED.published").
		Exec(c)
	return err
}
