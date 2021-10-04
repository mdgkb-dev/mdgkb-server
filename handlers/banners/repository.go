package banners

import (
	"fmt"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type IRepository interface {
	getAll(*gin.Context) ([]models.Banner, error)
	get(*gin.Context, string) (models.Banner, error)
	create(*gin.Context, *models.Banner) error
	delete(*gin.Context, string) error
	update(*gin.Context, *models.Banner) error
	updateAllOrder(*gin.Context, []*models.Banner) error
}

type Repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) getAll(ctx *gin.Context) (items []models.Banner, err error) {
	err = r.db.NewSelect().Model(&items).
		Relation("FileInfo").
		Order("list_number").
		Scan(ctx)
	return items, err
}

func (r *Repository) get(ctx *gin.Context, id string) (item models.Banner, err error) {
	err = r.db.NewSelect().Model(&item).Where("banner.id = ?", id).
	Relation("FileInfo").
	Scan(ctx)
	return item, err
}

func (r *Repository) create(ctx *gin.Context, item *models.Banner) (err error) {
	_, err = r.db.NewInsert().Model(item.FileInfo).Exec(ctx)
	item.FileInfoId = item.FileInfo.ID.UUID
	_, err = r.db.NewInsert().Model(item).Exec(ctx)
	fmt.Println(err)
	return err
}

func (r *Repository) delete(ctx *gin.Context, id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Banner{}).Where("id = ?", id).Exec(ctx)
	return err
}

func (r *Repository) update(ctx *gin.Context, item *models.Banner) (err error) {
	if item.FileInfo.ID.UUID != uuid.Nil {
		_, err = r.db.NewUpdate().Model(item.FileInfo).Where("id = ?", item.FileInfo.ID).Exec(ctx)
	}
	item.FileInfoId = item.FileInfo.ID.UUID
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(ctx)
	return err
}

func (r *Repository) updateAllOrder(ctx *gin.Context, item []*models.Banner) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
  Model(&item).
	Set("list_number = EXCLUDED.list_number").
	Where("banner.id = EXCLUDED.id").
	 Exec(ctx)
  return err
}
