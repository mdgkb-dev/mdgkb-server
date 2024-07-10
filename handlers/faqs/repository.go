package faqs

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"

	"github.com/uptrace/bun"
)

func (r *Repository) Create(c context.Context, item *models.Faq) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (models.Faqs, error) {
	items := make(models.Faqs, 0)
	err := r.helper.DB.IDB(c).NewSelect().Model(&items).Scan(c)
	return items, err
}

func (r *Repository) Get(c context.Context, id string) (*models.Faq, error) {
	item := models.Faq{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).Where("id = ?", id).Scan(c)
	return &item, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.Faq{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) UpsertMany(c context.Context, items models.Faqs) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Set("question = EXCLUDED.question").
		Set("answer = EXCLUDED.answer").
		Model(&items).
		Exec(c)
	return err
}

func (r *Repository) DeleteMany(c context.Context, idPool []uuid.UUID) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().
		Model((*models.Faq)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.Faq) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}
