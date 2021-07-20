package news

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

type Repository interface {
	create(*gin.Context, *models.News) error
	getAll(*gin.Context, *newsParams) ([]models.News, error)
	updateStatus(*gin.Context, *models.News) error
	delete(*gin.Context, string) error
}

type SRepository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *SRepository {
	return &SRepository{db}
}

func (r *SRepository) create(ctx *gin.Context, news *models.News) (err error) {
	_, err = r.db.NewInsert().Model(news).Exec(ctx)
	return err
}

func (r *SRepository) getAll(ctx *gin.Context, newsParams *newsParams) (news []models.News, err error) {
	query := r.db.NewSelect().Model(&news).Relation("Categories").Relation("Tags").Relation("PreviewThumbnailFile").Relation("NewsLikes")
	query = query.Order("published_on DESC").Limit(8)
	if newsParams.PublishedOn != nil {
		query = query.Where("published_on < ?", newsParams.PublishedOn)
	}
	err = query.Scan(ctx)
	return news, err
}

func (r *SRepository) updateStatus(ctx *gin.Context, news *models.News) (err error) {
	_, err = r.db.NewUpdate().Model(news).Exec(ctx)
	return err
}

func (r *SRepository) delete(ctx *gin.Context, id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.News{}).Where("id = ?", id).Exec(ctx)
	return err
}
