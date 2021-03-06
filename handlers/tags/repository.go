package tags

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
)

type IRepository interface {
	create(*gin.Context, *models.Tag) error
	getAll(*gin.Context) ([]models.Tag, error)
	get(*gin.Context, string) (models.Tag, error)
	updateStatus(*gin.Context, *models.Tag) error
	delete(*gin.Context, string) error
}

type Repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) create(ctx *gin.Context, item *models.Tag) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(ctx)
	return err
}

func (r *Repository) getAll(ctx *gin.Context) (items []models.Tag, err error) {
	err = r.db.NewSelect().Model(&items).Scan(ctx)
	return items, err
}

func (r *Repository) get(ctx *gin.Context, id string) (item models.Tag, err error) {
	err = r.db.NewSelect().Model(&item).Where("id = ?", id).Scan(ctx)
	return item, err
}

func (r *Repository) updateStatus(ctx *gin.Context, item *models.Tag) (err error) {
	_, err = r.db.NewUpdate().Model(item).Exec(ctx)
	return err
}

func (r *Repository) delete(ctx *gin.Context, id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Tag{}).Where("id = ?", id).Exec(ctx)
	return err
}
