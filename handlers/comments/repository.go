package comments

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items models.Comments) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.Comment)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Comments) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("user_id = EXCLUDED.user_id").
		Set("text = EXCLUDED.text").
		Set("rating = EXCLUDED.rating").
		Set("published_on = EXCLUDED.published_on").
		Exec(r.ctx)
	return err
}

func (r *Repository) get(id uuid.UUID) (item models.Comment, err error) {
	err = r.db().NewSelect().Model(&item).
		Relation("NewsComment.News").
		Relation("DoctorComment.Doctor.Human").
		Relation("DivisionComment.Division").
		Relation("User.Human").
		Where("comments.id = ?", id).Scan(r.ctx)

	return item, err
}

func (r *Repository) getAll() (item models.CommentsWithCount, err error) {
	item.Comments = make(models.Comments, 0)
	query := r.db().NewSelect().Model(&item.Comments).
		Relation("NewsComment.News").
		Relation("DoctorComment.Doctor.Human").
		Relation("DivisionComment.Division").
		Relation("User.Human")
	r.queryFilter.HandleQuery(query)
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) getAllMain() (models.Comments, error) {
	items := make(models.Comments, 0)
	query := r.db().NewSelect().Model(&items).Where("comments.positive = true").Order("published_on desc").Limit(4)
	err := query.Scan(r.ctx)

	return items, err
}

func (r *Repository) updateOne(item *models.Comment) error {
	_, err := r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertOne(item *models.Comment) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(item).
		Set("user_id = EXCLUDED.user_id").
		Set("text = EXCLUDED.text").
		Set("rating = EXCLUDED.rating").
		Set("published_on = EXCLUDED.published_on").
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
