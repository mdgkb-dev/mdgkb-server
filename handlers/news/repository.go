package news

import (
	"context"

	"mdgkb/mdgkb-server/models"
	"mdgkb/mdgkb-server/models/exportmodels"

	"github.com/uptrace/bun"
	//_ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) Create(c context.Context, news *models.News) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(news).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, news *models.News) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(news).Where("id = ?", news.ID).Exec(c)
	return err
}

func (r *Repository) CreateLike(c context.Context, item *models.NewsLike) error {
	_, err := r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) AddTag(c context.Context, item *models.NewsToTag) error {
	_, err := r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) RemoveTag(c context.Context, item *models.NewsToTag) error {
	_, err := r.helper.DB.IDB(c).NewDelete().Model(&models.NewsToTag{}).Where("news_id = ? AND tag_id = ?", item.NewsID, item.TagID).Exec(c)
	return err
}

func (r *Repository) CreateComment(c context.Context, item *models.NewsComment) error {
	_, err := r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) UpdateComment(c context.Context, item *models.NewsComment) error {
	_, err := r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}

func (r *Repository) RemoveComment(c context.Context, id string) error {
	_, err := r.helper.DB.IDB(c).NewDelete().Model(&models.NewsComment{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) GetAll(c context.Context) (items models.NewsWithCount, err error) {
	items.News = make([]*models.News, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&items.News).
		Relation("NewsToCategories.Category").
		Relation("NewsToTags.Tag").
		Relation("PreviewImage").
		Relation("NewsLikes").
		Relation("NewsViews")

	r.helper.SQL.ExtractFTSP(c).HandleQuery(query)
	items.Count, err = query.ScanAndCount(c)
	return items, err
}

func (r *Repository) GetMain(c context.Context) (items models.NewsWithCount, err error) {
	items.News = make([]*models.News, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&items.News).
		Relation("NewsToCategories.Category").
		Relation("NewsToTags.Tag").
		Relation("PreviewImage").
		Relation("NewsLikes").
		Relation("NewsViews").
		Where("news_view.main = ?", true)
	items.Count, err = query.ScanAndCount(c)
	return items, err
}

func (r *Repository) GetSubMain(c context.Context) (items models.NewsWithCount, err error) {
	items.News = make([]*models.News, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&items.News).
		Relation("NewsToCategories.Category").
		Relation("NewsToTags.Tag").
		Relation("PreviewImage").
		Relation("NewsLikes").
		Relation("NewsViews").
		Where("news_view.sub_main = ?", true)
	items.Count, err = query.ScanAndCount(c)
	return items, err
}

func (r *Repository) GetBySlug(c context.Context, slug string) (*models.News, error) {
	item := models.News{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("NewsToCategories.Category").
		Relation("NewsToTags.Tag").
		Relation("PreviewImage").
		Relation("MainImage").
		Relation("NewsLikes").
		Relation("NewsViews").
		Relation("Event.Form.Fields.ValueType").
		Relation("Event.EventApplications.FieldValues").
		Relation("Event.EventApplications.User").
		// Relation("NewsComments.Comment", func(q *bun.SelectQuery) *bun.SelectQuery {
		// 	return q.Where("comment.mod_checked = true").Order("comment.published_on DESC")
		// }).
		// Relation("NewsComments.Comment.User.Human").
		Relation("NewsImages", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("news_images.news_image_order")
		}).
		Relation("NewsImages.FileInfo").
		Relation("NewsDoctors.Doctor").
		Relation("NewsDoctors.Doctor.Employee.Human").
		Relation("NewsDoctors.Doctor.Employee.Human.PhotoMini").
		Relation("NewsDoctors.Doctor.Employee.Regalias").
		Where("news_view.id = ?", slug).Scan(c)
	return &item, err
}

func (r *Repository) GetNewsComments(c context.Context, id string) (*models.NewsComments, error) {
	items := models.NewsComments{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&items).
		Relation("Comment", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("comment.mod_checked = true").Order("comment.published_on DESC")
		}).
		Relation("Comment.User.Human").
		Where("news_comments.news_id = ?", id).Scan(c)
	return &items, err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.News{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) DeleteLike(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.NewsLike{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) CreateViewOfNews(c context.Context, newsView *models.NewsView) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(newsView).On("CONFLICT (ip_address, news_id) DO NOTHING").Exec(c)
	return err
}

func (r *Repository) GetSuggestionNews(c context.Context, id string) ([]*models.News, error) {
	items := make([]*models.News, 0)
	generalNewsQuery := r.helper.DB.IDB(c).NewSelect().
		ColumnExpr("news_view.id as news_id, news_to_tags.tag_id as tag_id").
		Model((*models.News)(nil)).
		Join("join news_to_tags on news_to_tags.news_id = news_view.id").
		Where("?TableAlias.id =? ", id)

	err := r.helper.DB.IDB(c).NewSelect().
		With("gen_news", generalNewsQuery).
		Model(&items).
		ColumnExpr("news_view.id, news_view.title, news_view.published_on,views_count, count(news_to_tags.id) as tag_count ").
		Join("left join news_to_tags on news_to_tags.news_id = news_view.id").
		Join("left join gen_news on gen_news.tag_id = news_to_tags.tag_id").
		Group("news_view.id", "news_to_tags.tag_id", "news_view.title", "news_view.published_on", "views_count").
		Order("tag_count desc", "views_count", "published_on desc").
		Limit(4).
		Scan(c)
	return items, err
}

func (r *Repository) GetAggregateViews(c context.Context, opt *exportmodels.NewsView) (models.ChartDataSets, error) {
	items := make(models.ChartDataSets, 0)
	query := r.helper.DB.IDB(c).NewSelect().
		ColumnExpr(opt.GetColExpr()).
		TableExpr("news n").
		Join("join news_views nv on n.id = nv.news_id and nv.created_at is not null").
		Group(opt.GetGroupExpr())

	if len(opt.IDPool) > 0 {
		query.Where("n.id in (?)", bun.In(opt.IDPool))
	}

	err := query.Scan(c, &items)
	return items, err
}
