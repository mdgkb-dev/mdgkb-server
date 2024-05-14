package comments

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"

	"github.com/uptrace/bun"
)

func (r *Repository) CreateMany(c context.Context, items models.Comments) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(&items).Exec(c)
	return err
}

func (r *Repository) DeleteMany(c context.Context, idPool []string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().
		Model((*models.Comment)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(c)
	return err
}

func (r *Repository) UpsertMany(c context.Context, items models.Comments) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("user_id = EXCLUDED.user_id").
		Set("text = EXCLUDED.text").
		Set("rating = EXCLUDED.rating").
		Set("published_on = EXCLUDED.published_on").
		Exec(c)
	return err
}

func (r *Repository) Create(c context.Context, item *models.Comment) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(&item).Exec(c)
	return err
}

func (r *Repository) Get(c context.Context, id uuid.NullUUID) (item models.Comment, err error) {
	err = r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("NewsComment.News").
		Relation("DoctorComment.Doctor.Employee.Human").
		Relation("DivisionComment.Division").
		Relation("User.Human").
		Where("comments.id = ?", id).Scan(c)

	return item, err
}

func (r *Repository) GetAll(c context.Context) (item models.CommentsWithCount, err error) {
	item.Comments = make(models.Comments, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&item.Comments).
		Relation("NewsComment.News").
		Relation("DoctorComment.Doctor.Employee.Human").
		Relation("DivisionComment.Division").
		Relation("User.Human")

	r.helper.SQL.ExtractFTSP(c).HandleQuery(query)
	item.Count, err = query.ScanAndCount(c)
	return item, err
}

func (r *Repository) GetAllMain(c context.Context) (models.Comments, error) {
	items := make(models.Comments, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&items).Where("comments.positive = true").Order("published_on desc").Limit(4)
	err := query.Scan(c)

	return items, err
}

func (r *Repository) UpdateOne(c context.Context, item *models.Comment) error {
	_, err := r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) UpsertOne(c context.Context, item *models.Comment) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("conflict (id) do update").
		Model(item).
		Set("user_id = EXCLUDED.user_id").
		Set("text = EXCLUDED.text").
		Set("rating = EXCLUDED.rating").
		Set("published_on = EXCLUDED.published_on").
		Exec(c)
	return err
}
