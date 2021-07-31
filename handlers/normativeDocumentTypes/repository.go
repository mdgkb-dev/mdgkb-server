package normativeDocumentTypes

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type IRepository interface {
	create(*gin.Context, *models.NormativeDocumentType) error
	get(*gin.Context, string) (models.NormativeDocumentType, error)
	getAll(*gin.Context) ([]models.NormativeDocumentType, error)
	update(*gin.Context, *models.NormativeDocumentType) error
	delete(*gin.Context, string) error
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

func (r *Repository) get(ctx *gin.Context, id string) (item models.NormativeDocumentType, err error) {
	err = r.db.NewSelect().Model(&item).Where("id = ?", id).Scan(ctx)
	return item, err
}

func (r *Repository) getAll(ctx *gin.Context) (items []models.NormativeDocumentType, err error) {
	err = r.db.NewSelect().Model(&items).Scan(ctx)
	return items, err
}

func (r *Repository) update(ctx *gin.Context, documentType *models.NormativeDocumentType) (err error) {
	_, err = r.db.NewUpdate().Model(documentType).Where("id = ?", documentType.ID).Exec(ctx)
	return err
}

func (r *Repository) delete(ctx *gin.Context, id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.NormativeDocumentType{}).Where("id = ?", id).Exec(ctx)
	return err
}
