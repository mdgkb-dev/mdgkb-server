package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PathPermissionRole struct {
	bun.BaseModel    `bun:"path_permissions_roles,alias:path_permissions_roles"`
	ID               uuid.NullUUID   `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	PathPermission   *PathPermission `bun:"rel:belongs-to" json:"pathPermission"`
	PathPermissionID uuid.NullUUID   `bun:"type:uuid" json:"pathPermissionId"`
	Role             *Role           `bun:"rel:belongs-to" json:"role"`
	RoleID           uuid.NullUUID   `bun:"type:uuid" json:"roleId"`
}

type PathPermissionsRoles []*PathPermissionRole
