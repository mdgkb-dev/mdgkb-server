package normativeDocumentTypes

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IRepository interface {
	create(*gin.Context, *models.NormativeDocumentType) error
}

type Repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) create(ctx *gin.Context, item *models.NormativeDocumentType) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(ctx)
	return err
}
