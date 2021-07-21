package news

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

type IRepository interface {
	create(*gin.Context, *models.News) error
	createLike(*gin.Context, *models.NewsLike) error
	getAll(*gin.Context, *newsParams) ([]models.News, error)
	updateStatus(*gin.Context, *models.News) error
	delete(*gin.Context, string) error
	deleteLike(*gin.Context, string) error
	getBySlug(*gin.Context, string) (models.News, error)
}

type Repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) create(ctx *gin.Context, news *models.News) (err error) {
	_, err = r.db.NewInsert().Model(news).Exec(ctx)
	return err
}

func (r *Repository) createLike(ctx *gin.Context, item *models.NewsLike) error {
	_, err := r.db.NewInsert().Model(item).Exec(ctx)
	return err
}

func (r *Repository) getAll(ctx *gin.Context, newsParams *newsParams) (news []models.News, err error) {
	query := r.db.NewSelect().Model(&news).Relation("Categories").Relation("Tags").Relation("PreviewThumbnailFile").Relation("NewsLikes")
	query = query.Order("published_on DESC").Limit(6)
	if newsParams.PublishedOn != nil {
		query = query.Where("published_on < ?", newsParams.PublishedOn)
	}
	err = query.Scan(ctx)
	return news, err
}

func (r *Repository) getBySlug(ctx *gin.Context, slug string) (item models.News, err error) {
	err = r.db.NewSelect().Model(&item).Where("slug = ?", slug).Scan(ctx)
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
