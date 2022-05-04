package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) upsertManyPathPermissions(items models.PathPermissions) (err error) {
	_, err = r.db.NewInsert().On("CONFLICT (resource) DO UPDATE").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("guest_allow = EXCLUDED.guest_allow").
		Exec(r.ctx)
	fmt.Println("ERRRRRRRRRRRRRRRRRRRRRRR", err)
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
	_, err = r.db.NewInsert().On("CONFLICT (path_permission_id, role_id) DO UPDATE").
		Set("path_permission_id = EXCLUDED.path_permission_id").
		Set("role_id = EXCLUDED.role_id").
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

func (r *Repository) getAllPathPermissionsAdmin() (items models.PathPermissionsWithCount, err error) {
	query := r.db.NewSelect().Model(&items.PathPermissions).
		Relation("PathPermissionsRoles")

	r.queryFilter.Paginator.CreatePagination(query)
	r.queryFilter.Filter.CreateFilter(query)
	r.queryFilter.Sorter.CreateOrder(query)
	items.Count, err = query.ScanAndCount(r.ctx)
	return items, err
}

func (r *Repository) checkPathPermissions(path string, roleID string) error {
	if roleID == "" {
		return r.db.NewSelect().
			Model(&models.PathPermission{}).
			Where("path_permissions.resource = ? and path_permissions.guest_allow = true", path).
			Scan(r.ctx)
	}

	return r.db.NewSelect().
		Model(&models.PathPermission{}).
		Join("JOIN path_permissions_roles ppr on ppr.path_permission_id = path_permissions.id and ppr.role_id = ?", roleID).
		Where("path_permissions.resource = ?", path).
		Scan(r.ctx)

}

func (r *Repository) getPathPermissionsByRoleId(id string) (models.PathPermissions, error) {
	items := make(models.PathPermissions, 0)
	err := r.db.NewSelect().
		Model(&items).
		Relation("PathPermissionsRoles").
		Join("JOIN path_permissions_roles ppr on ppr.path_permission_id = path_permissions.id and ppr.role_id = ?", id).
		// Where("path_permissions.path_permissions_roles.role_id = ?", id).
		Scan(r.ctx)
	return items, err
}
