package news

import (
	"mdgkb/mdgkb-server/models"
	"strings"

	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(news *models.News) (err error) {
	_, err = r.db.NewInsert().Model(news).Exec(r.ctx)
	return err
}

func (r *Repository) update(news *models.News) (err error) {
	_, err = r.db.NewUpdate().Model(news).Where("id = ?", news.ID).Exec(r.ctx)
	return err
}

func (r *Repository) createLike(item *models.NewsLike) error {
	_, err := r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) addTag(item *models.NewsToTag) error {
	_, err := r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) removeTag(item *models.NewsToTag) error {
	_, err := r.db.NewDelete().Model(&models.NewsToTag{}).Where("news_id = ? AND tag_id = ?", item.NewsID, item.TagID).Exec(r.ctx)
	return err
}

func (r *Repository) createComment(item *models.NewsComment) error {
	_, err := r.db.NewInsert().Model(item.Comment).Exec(r.ctx)
	item.CommentID = item.Comment.ID
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) updateComment(item *models.NewsComment) error {
	_, err := r.db.NewUpdate().Model(item.Comment).Where("id = ?", item.Comment.ID).Exec(r.ctx)
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) removeComment(id string) error {
	_, err := r.db.NewDelete().Model(&models.NewsComment{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) getAll(newsParams *newsParams) (news []models.News, err error) {
	query := r.db.NewSelect().Model(&news).
		Relation("NewsToCategories.Category").
		Relation("NewsToTags.Tag").
		Relation("FileInfo").
		Relation("NewsLikes").
		Relation("NewsViews")

	if newsParams.Limit != 0 {
		query = query.Order("published_on DESC").Limit(newsParams.Limit)
	}
	if newsParams.PublishedOn != nil {
		query = query.Where("published_on < ?", newsParams.PublishedOn)
	}
	if newsParams.FilterTags != "" && newsParams.OrderByView == "" {
		for _, tagId := range strings.Split(newsParams.FilterTags, ",") {
			query = query.Where("exists (select * from news_to_tags as ntt where ntt.news_id = news.id and ntt.tag_id = ?)", tagId)
		}
	}
	if newsParams.FilterTags != "" && newsParams.OrderByView != "" && newsParams.Limit != 0 {
		query = query.Join("JOIN news_to_tags ON news_to_tags.news_id = news.id").
			Join("LEFT JOIN news_views ON news_views.news_id = news.id").
			Where("news_to_tags.tag_id in (?)", bun.In(strings.Split(newsParams.FilterTags, ","))).
			Group("news.id", "file_info.id").
			OrderExpr("count (news_to_tags.id)").
			OrderExpr("count (news_views.id)").
			Limit(newsParams.Limit)
		if len(news) == 0 {
			query = r.db.NewSelect().Model(&news).Relation("NewsViews").
				Join("LEFT JOIN news_views ON news_views.news_id = news.id").
				Group("news.id").
				Order("published_on DESC").
				OrderExpr("count (news_views.id)").
				Limit(newsParams.Limit)
		}
	}
	if newsParams.Events {
		query = query.Join("JOIN events ON events.id = news.event_id")
	}
	err = query.Scan(r.ctx)
	return news, err
}

func (r *Repository) getBySlug(slug string) (*models.News, error) {
	item := new(models.News)
	err := r.db.NewSelect().Model(item).
		Relation("NewsToCategories.Category").
		Relation("NewsToTags.Tag").
		Relation("FileInfo").
		Relation("MainImage").
		Relation("NewsLikes").
		Relation("NewsViews").
		Relation("Event").
		Relation("NewsComments.Comment.User").
		Relation("NewsImages.FileInfo").
		Where("news.slug = ?", slug).Scan(r.ctx)
	return item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.News{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) deleteLike(id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.NewsLike{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) getByMonth(monthParams *monthParams) (news []models.News, err error) {
	query := r.db.NewSelect().Model(&news)
	query = query.Where("extract(year from news.published_on) = ?", monthParams.Year).Where("extract(month from news.published_on) = ?", monthParams.Month)
	err = query.Scan(r.ctx)
	return news, err
}

func (r *Repository) createViewOfNews(newsView *models.NewsView) (err error) {
	_, err = r.db.NewInsert().Model(newsView).On("CONFLICT (ip_address, news_id) DO NOTHING").Exec(r.ctx)
	return err
}

func (r *Repository) createEventApplication(item *models.EventApplication) error {
	_, err := r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}
