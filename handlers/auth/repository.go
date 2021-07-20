package auth

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

type Repository interface {
	getByLogin(*gin.Context, *string) (*models.User, error)
	create(*gin.Context, *models.User) error
}

type ARepository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *ARepository {
	return &ARepository{db}
}

func (r *ARepository) getByLogin(ctx *gin.Context, email *string) (user *models.User, err error) {
	err = r.db.NewSelect().Model(&user).
		Where("email = ?", *email).
		Scan(ctx)
	return user, err
}

func (r *ARepository) create(ctx *gin.Context, user *models.User) (err error) {
	_, err = r.db.NewInsert().Model(user).Exec(ctx)
	return err
}
