package entrances

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) getAll() (models.Entrances, error) {
	items := make(models.Entrances, 0)
	err := r.db.NewSelect().Model(&items).Relation("Building").Scan(r.ctx)
	return items, err
}
