package users

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

type IRepository interface {
	getAll(*gin.Context) ([]models.User, error)
	get(*gin.Context, string) (models.User, error)
	getByEmail(*gin.Context, string) (models.User, error)
}

type Repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) getAll(ctx *gin.Context) (items []models.User, err error) {
	err = r.db.NewSelect().Model(&items).Scan(ctx)
	return items, err
}

func (r *Repository) get(ctx *gin.Context, id string) (item models.User, err error) {
	err = r.db.NewSelect().Model(&item).Where("id = ?", id).Scan(ctx)
	return item, err
}

func (r *Repository) getByEmail(ctx *gin.Context, id string) (item models.User, err error) {
	err = r.db.NewSelect().Model(&item).Where("email = ?", id).Scan(ctx)
	return item, err
}
