package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PathPermission struct {
	bun.BaseModel        `bun:"path_permissions,alias:path_permissions"`
	ID                   uuid.UUID            `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Resource             string               `json:"resource"`
	PathPermissionsRoles PathPermissionsRoles `bun:"rel:has-many" json:"pathPermissionsRoles"`
}

type PathPermissions []*PathPermission
