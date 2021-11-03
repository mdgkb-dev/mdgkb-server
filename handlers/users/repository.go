package users

import (
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"

	_ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) getAll() (models.Users, error) {
	items := make(models.Users, 0)
	err := r.db.NewSelect().Model(&items).Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.User, error) {
	item := models.User{}
	err := r.db.NewSelect().
		Model(&item).Relation("Human").
		Where("users.id = ?", id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) getByEmail(id string) (*models.User, error) {
	item := models.User{}
	err := r.db.NewSelect().Model(&item).Where("users.email = ?", id).Scan(r.ctx)
	return &item, err
}
