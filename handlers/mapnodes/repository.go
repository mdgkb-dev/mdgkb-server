package mapnodes

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	// _ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UploadMapNodes(items models.MapNodes) (err error) {
	_, err = r.db().NewInsert().Model(items).Exec(r.ctx)
	return err
}

func (r *Repository) DeleteAll() error {
	_, err := r.db().NewDelete().Model(&models.MapNode{}).Where("id is not null").Exec(r.ctx)
	return err
}

func (r *Repository) CreateMany(items models.MapNodes) error {
	_, err := r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}
