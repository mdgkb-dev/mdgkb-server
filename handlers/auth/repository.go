package auth

import (
	"fmt"
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

func (r *Repository) checkPathPermissions(path string, roleID string) error {
	err := r.db.NewSelect().
		Model(&models.PathPermission{}).
		Join("JOIN path_permissions_roles ppr on ppr.path_permission_id = path_permissions.id and ppr.role_id = ?", roleID).
		Where("path_permissions.resource = ?", path).
		Scan(r.ctx)
	fmt.Println(err)
	return err
}
