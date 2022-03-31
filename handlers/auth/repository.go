package auth

import (
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) savePathPermissions(items models.PathPermissions) (err error) {
	_, err = r.db.NewInsert().
		Model(&items).
		Exec(r.ctx)
	return err
}

func (r *Repository) getAllPathPermissions() (models.PathPermissions, error) {
	items := make(models.PathPermissions, 0)
	err := r.db.NewSelect().Model(&items).
		Scan(r.ctx)
	return items, err
}
