package news

import (
	"mdgkb/mdgkb-server/models"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/google/uuid"

	"github.com/uptrace/bun"
)

type IRepository interface {
	create(*gin.Context, *models.News) error
	update(*gin.Context, *models.News) error
	createLike(*gin.Context, *models.NewsLike) error
	addTag(*gin.Context, *models.NewsToTag) error
	removeTag(*gin.Context, *models.NewsToTag) error
	removeComment(*gin.Context, string) error
	createComment(*gin.Context, *models.NewsComment) error
	updateComment(*gin.Context, *models.NewsComment) error
	getAll(*gin.Context, *newsParams) ([]models.News, error)
	updateStatus(*gin.Context, *models.News) error
	delete(*gin.Context, string) error
	deleteLike(*gin.Context, string) error
	getBySlug(*gin.Context, string) (models.News, error)
	getByMonth(*gin.Context, *monthParams) ([]models.News, error)
}

type Repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) create(ctx *gin.Context, news *models.News) (err error) {
	_, err = r.db.NewInsert().Model(news.FileInfo).Exec(ctx)
	news.FileInfoId = news.FileInfo.ID
	_, err = r.db.NewInsert().Model(news).Exec(ctx)

	for _, tag := range news.Tags {
		newsTags := models.NewsToTag{TagId: tag.ID, NewsId: news.ID}
		_, err = r.db.NewInsert().Model(newsTags).Exec(ctx)
	}

	return err
}

func (r *Repository) update(ctx *gin.Context, news *models.News) (err error) {
	if news.FileInfo.ID != uuid.Nil {
		_, err = r.db.NewUpdate().Model(news.FileInfo).Where("id = ?", news.FileInfo.ID).Exec(ctx)
	}
	news.FileInfoId = news.FileInfo.ID
	_, err = r.db.NewUpdate().Model(news).Where("id = ?", news.ID).Exec(ctx)

	// TODO Стас, посмотри, плз
	newsTags := new([]models.NewsToTag)
	r.db.NewSelect().Model(newsTags).Where("news_id = ?", news.ID).Scan(ctx)
	for j := 0; j < len(*newsTags); j++ {
		found := false
		for i := 0; i < len(news.Tags); i++ {
			if news.Tags[i].ID == (*newsTags)[j].ID {
				found = true
			}
		}
		if !found {
			_, err = r.db.NewDelete().Model(newsTags).Where("id = ?", (*newsTags)[j].ID).Exec(ctx)
		}
	}

	for _, tag := range news.Tags {
		newsTag := new(models.NewsToTag)
		r.db.NewSelect().Model(newsTag).Where("tag_id = ?", tag.ID).Scan(ctx)
		newsTags := models.NewsToTag{TagId: tag.ID, NewsId: news.ID}
		_, err = r.db.NewInsert().Model(&newsTags).Exec(ctx)
	}

	return err
}

func (r *Repository) createLike(ctx *gin.Context, item *models.NewsLike) error {
	_, err := r.db.NewInsert().Model(item).Exec(ctx)
	return err
}

func (r *Repository) addTag(ctx *gin.Context, item *models.NewsToTag) error {
	_, err := r.db.NewInsert().Model(item).Exec(ctx)
	return err
}

func (r *Repository) removeTag(ctx *gin.Context, item *models.NewsToTag) error {
	_, err := r.db.NewDelete().Model(&models.NewsToTag{}).Where("news_id = ? AND tag_id = ?", item.NewsId, item.TagId).Exec(ctx)
	return err
}

func (r *Repository) removeComment(ctx *gin.Context, id string) error {
	_, err := r.db.NewDelete().Model(&models.NewsComment{}).Where("id = ?", id).Exec(ctx)
	return err
}

func (r *Repository) createComment(ctx *gin.Context, item *models.NewsComment) error {
	_, err := r.db.NewInsert().Model(item).Exec(ctx)
	return err
}

func (r *Repository) updateComment(ctx *gin.Context, item *models.NewsComment) error {
	_, err := r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(ctx)
	return err
}

func (r *Repository) getAll(ctx *gin.Context, newsParams *newsParams) (news []models.News, err error) {
	query := r.db.NewSelect().Model(&news).
		Relation("Categories").
		Relation("Tags").
		Relation("FileInfo").
		Relation("NewsLikes")

	if newsParams.Limit != 0 {
		query = query.Order("published_on DESC").Limit(newsParams.Limit)
	}
	if newsParams.PublishedOn != nil {
		query = query.Where("published_on < ?", newsParams.PublishedOn)
	}
	if newsParams.FilterTags != "" {
		for _, tagId := range strings.Split(newsParams.FilterTags, ",") {
			query = query.Where("exists (select * from news_to_tags as ntt where ntt.news_id = news.id and ntt.tag_id = ?)", tagId)
		}
	}
	err = query.Scan(ctx)
	return news, err
}

func (r *Repository) getBySlug(ctx *gin.Context, slug string) (item models.News, err error) {
	err = r.db.NewSelect().Model(&item).
		Relation("Categories").
		Relation("Tags").
		Relation("FileInfo").
		Relation("NewsLikes").
		Relation("NewsComments.User").
		Relation("NewsImages.FileInfo").
		Where("slug = ?", slug).Scan(ctx)
	return item, err
}

func (r *Repository) updateStatus(ctx *gin.Context, news *models.News) (err error) {
	_, err = r.db.NewUpdate().Model(news).Exec(ctx)
	return err
}

func (r *Repository) delete(ctx *gin.Context, id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.News{}).Where("id = ?", id).Exec(ctx)
	return err
}

func (r *Repository) deleteLike(ctx *gin.Context, id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.NewsLike{}).Where("id = ?", id).Exec(ctx)
	return err
}

func (r *Repository) getByMonth(ctx *gin.Context, monthParams *monthParams) (news []models.News, err error) {
	query := r.db.NewSelect().Model(&news)
	query = query.Where("extract(year from news.published_on) = ?", monthParams.Year).Where("extract(month from news.published_on) = ?", monthParams.Month)
	err = query.Scan(ctx)
	return news, err
}
