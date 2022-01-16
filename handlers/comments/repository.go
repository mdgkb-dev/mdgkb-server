package comments

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) createMany(items models.Comments) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.Comment)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Comments) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("user_id = EXCLUDED.user_id").
		Set("text = EXCLUDED.user_id").
		Set("rating = EXCLUDED.rating").
		Set("published_on = EXCLUDED.published_on").
		Exec(r.ctx)
	return err
}

func (r *Repository) getAll(params *commentsParams) (models.Comments, error) {
	items := make(models.Comments, 0)
	query := r.db.NewSelect().Model(&items).
		Relation("NewsComment.News").
		Relation("DoctorComment.Doctor.Human").
		Relation("DivisionComment.Division").
		Relation("User").
		Order("published_on DESC")
	if params.Limit != 0 {
		query = query.Limit(params.Limit)
	}
	if params.ModChecked != nil {
		query = query.Where("comment.mod_checked = ?", params.ModChecked)
	}
	if params.Positive != nil {
		query = query.Where("comment.positive = ?", params.Positive)
	}
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) updateOne(item *models.Comment) error {
	_, err := r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
