package schedules

import (
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"

	_ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Schedule) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}


func (r *Repository) upsert(item *models.Schedule) (err error) {
	_, err = r.db.NewInsert().Model(item).On("conflict (id) do update").
		Set("name = EXCLUDED.name").
		Set("description = EXCLUDED.description").
		Exec(r.ctx)
	return err
}
