package tags

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(ctx *gin.Context, item *models.Tag) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(ctx)
	return err
}

func (r *Repository) getAll(ctx *gin.Context) (items []models.Tag, err error) {
	err = r.db().NewSelect().Model(&items).Scan(ctx)
	return items, err
}

func (r *Repository) get(ctx *gin.Context, id string) (item models.Tag, err error) {
	err = r.db().NewSelect().Model(&item).Where("id = ?", id).Scan(ctx)
	return item, err
}

func (r *Repository) updateStatus(ctx *gin.Context, item *models.Tag) (err error) {
	_, err = r.db().NewUpdate().Model(item).Exec(ctx)
	return err
}

func (r *Repository) delete(ctx *gin.Context, id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Tag{}).Where("id = ?", id).Exec(ctx)
	return err
}
