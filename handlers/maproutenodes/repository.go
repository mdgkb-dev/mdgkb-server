package maproutenodes

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
	// _ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) DeleteAll() error {
	_, err := r.db().NewDelete().Model(&models.MapRouteNode{}).Where("map_node_name is not null").Exec(r.ctx)
	return err
}

func (r *Repository) CreateMany(items models.MapRouteNodes) error {
	_, err := r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}
