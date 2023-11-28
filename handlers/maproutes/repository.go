package maproutes

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
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetMapRoute(startNodeID string, endNodeID string) (*models.MapRoute, error) {
	item := models.MapRoute{}
	err := r.db().NewSelect().Model(&item).
		Where("?TableAlias.start_node_id = ?", startNodeID).
		Where("?TableAlias.end_node_id = ?", endNodeID).
		Scan(r.ctx)

	return &item, err
}
