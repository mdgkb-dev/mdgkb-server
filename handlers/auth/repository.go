package auth

import (
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) upsertManyPathPermissions(items models.PathPermissions) (err error) {
	_, err = r.db.NewInsert().On("CONFLICT DO NOTHING").
		Model(&items).
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteManyPathPermissions(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.PathPermission)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertManyPathPermissionsRoles(items models.PathPermissionsRoles) (err error) {
	_, err = r.db.NewInsert().On("CONFLICT DO NOTHING").
		Model(&items).
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteManyPathPermissionsRoles(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.PathPermissionRole)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) getAllPathPermissions() (models.PathPermissions, error) {
	items := make(models.PathPermissions, 0)
	err := r.db.NewSelect().
		Model(&items).
		Relation("PathPermissionsRoles").
		Scan(r.ctx)
	return items, err
}
