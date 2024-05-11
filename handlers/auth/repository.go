package auth

import (
	"context"
	"mdgkb/mdgkb-server/models"

	// _ "github.com/go-pg/pg/v10/orm"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) UpsertManyPathPermissions(c context.Context, items models.PathPermissions) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("CONFLICT (resource) DO UPDATE").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("guest_allow = EXCLUDED.guest_allow").
		Exec(c)
	return err
}

func (r *Repository) DeleteManyPathPermissions(c context.Context, idPool []uuid.UUID) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().
		Model((*models.PathPermission)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(c)
	return err
}

func (r *Repository) UpsertManyPathPermissionsRoles(c context.Context, items models.PathPermissionsRoles) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().On("CONFLICT (path_permission_id, role_id) DO UPDATE").
		Set("path_permission_id = EXCLUDED.path_permission_id").
		Set("role_id = EXCLUDED.role_id").
		Model(&items).
		Exec(c)
	return err
}

func (r *Repository) DeleteManyPathPermissionsRoles(c context.Context, idPool []uuid.UUID) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().
		Model((*models.PathPermissionRole)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(c)
	return err
}

func (r *Repository) GetAllPathPermissions(c context.Context) (models.PathPermissions, error) {
	items := make(models.PathPermissions, 0)
	err := r.helper.DB.IDB(c).NewSelect().
		Model(&items).
		Relation("PathPermissionsRoles").
		Scan(c)
	return items, err
}

func (r *Repository) GetAllPathPermissionsAdmin(c context.Context) (item models.PathPermissionsWithCount, err error) {
	item.PathPermissions = make(models.PathPermissions, 0)
	query := r.helper.DB.IDB(c).NewSelect().Model(&item.PathPermissions).
		Relation("PathPermissionsRoles")

	item.Count, err = query.ScanAndCount(c)
	return item, err
}

func (r *Repository) CheckPathPermissions(c context.Context, path string, roleID string) error {
	if roleID == "" {
		return r.helper.DB.IDB(c).NewSelect().
			Model(&models.PathPermission{}).
			Where("path_permissions.resource = ? and path_permissions.guest_allow = true", path).
			Scan(c)
	}

	return r.helper.DB.IDB(c).NewSelect().
		Model(&models.PathPermission{}).
		Join("JOIN path_permissions_roles ppr on ppr.path_permission_id = path_permissions.id and ppr.role_id = ?", roleID).
		Where("path_permissions.resource = ?", path).
		Scan(c)
}

func (r *Repository) GetPathPermissionsByRoleID(c context.Context, id string) (models.PathPermissions, error) {
	items := make(models.PathPermissions, 0)
	err := r.helper.DB.IDB(c).NewSelect().
		Model(&items).
		Relation("PathPermissionsRoles").
		Join("JOIN path_permissions_roles ppr on ppr.path_permission_id = path_permissions.id and ppr.role_id = ?", id).
		// Where("path_permissions.path_permissions_roles.role_id = ?", id).
		Scan(c)
	return items, err
}
