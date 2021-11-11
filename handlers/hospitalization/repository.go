package hospitalization

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) getAll() (models.Hospitalizations, error) {
	items := make(models.Hospitalizations, 0)
	err := r.db.NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}
