package carousels

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
)

type IRepository interface {
	create(*gin.Context, *models.Carousel) error
	getAll(*gin.Context) ([]models.Carousel, error)
	get(*gin.Context, string) (models.Carousel, error)
	getByKey(*gin.Context, string) (models.Carousel, error)
	update(*gin.Context, *models.Carousel) error
	delete(*gin.Context, string) error
}

type Repository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) create(ctx *gin.Context, item *models.Carousel) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(ctx)

	if len(item.CarouselSlides) == 0 {
		return err
	}
	var fileInfos []models.FileInfo
	for _, slide := range item.CarouselSlides {
		fileInfos = append(fileInfos, *slide.FileInfo)
	}
	_, err = r.db.NewInsert().Model(&fileInfos).Exec(ctx)

	for i, slide := range item.CarouselSlides {
		slide.FileInfoId = fileInfos[i].ID.UUID
		slide.CarouselID = item.ID
	}

	_, err = r.db.NewInsert().Model(&item.CarouselSlides).Exec(ctx)

	return err
}

func (r *Repository) getAll(ctx *gin.Context) (items []models.Carousel, err error) {
	err = r.db.NewSelect().Model(&items).Scan(ctx)
	return items, err
}

func (r *Repository) get(ctx *gin.Context, id string) (item models.Carousel, err error) {
	err = r.db.NewSelect().
		Model(&item).Relation("CarouselSlides.FileInfo").Where("id = ?", id).Scan(ctx)
	return item, err
}

func (r *Repository) getByKey(ctx *gin.Context, key string) (item models.Carousel, err error) {
	err = r.db.NewSelect().
		Model(&item).Relation("CarouselSlides.FileInfo").Where("system_key = ?", key).Scan(ctx)
	return item, err
}

func (r *Repository) update(ctx *gin.Context, item *models.Carousel) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(ctx)

	if len(item.CarouselSlidesForDelete) > 0 {
		_, err = r.db.NewDelete().Model((*models.CarouselSlide)(nil)).Where("id IN (?)", bun.In(item.CarouselSlidesForDelete)).Exec(ctx)
	}

	if len(item.CarouselSlides) == 0 {
		return err
	}
	var fileInfos []models.FileInfo
	for _, slide := range item.CarouselSlides {
		fileInfos = append(fileInfos, *slide.FileInfo)
	}
	_, err = r.db.NewInsert().Model(&fileInfos).
		On("CONFLICT (id) DO UPDATE").
		Set("original_name = EXCLUDED.original_name").
		Set("file_system_path = EXCLUDED.file_system_path").
		Exec(ctx)

	for i, slide := range item.CarouselSlides {
		slide.FileInfoId = fileInfos[i].ID.UUID
		slide.CarouselID = item.ID
	}

	_, err = r.db.NewInsert().Model(&item.CarouselSlides).
		On("CONFLICT (id) DO UPDATE").
		Set("title = EXCLUDED.title").
		Set("content = EXCLUDED.content").
		Set("link = EXCLUDED.link").
		Set("button_show = EXCLUDED.button_show").
		Set("button_color = EXCLUDED.button_color").
		Exec(ctx)

	return err
}

func (r *Repository) delete(ctx *gin.Context, id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Carousel{}).Where("id = ?", id).Exec(ctx)
	return err
}
