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
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetMapRoute(startNodeID string, endNodeID string) (*models.MapRoute, error) {
	item := models.MapRoute{}
	err := r.db().NewSelect().Model(&item).
		Relation("MapRouteNodes", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("item_order")
		}).
		Where("?TableAlias.start_node_name = ? and ?TableAlias.end_node_name = ?", startNodeID, endNodeID).
		WhereOr("?TableAlias.start_node_name = ? and ?TableAlias.end_node_name = ?", endNodeID, startNodeID).
		Scan(r.ctx)

	return &item, err
}

func (r *Repository) DeleteAll() error {
	_, err := r.db().NewDelete().Model(&models.MapRoute{}).Where("id is not null").Exec(r.ctx)
	return err
}

func (r *Repository) CreateMany(items models.MapRoutes) error {
	_, err := r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}
