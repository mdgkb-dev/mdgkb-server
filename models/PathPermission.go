package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PathPermission struct {
	bun.BaseModel                 `bun:"path_permissions,alias:path_permissions"`
	ID                            uuid.NullUUID        `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Resource                      string               `bun:",unique" json:"resource"`
	GuestAllow                    bool                 `json:"guestAllow"`
	PathPermissionsRoles          PathPermissionsRoles `bun:"rel:has-many" json:"pathPermissionsRoles"`
	PathPermissionsRolesForDelete []uuid.UUID          `bun:"-" json:"pathPermissionsRolesForDelete"`

	//SortColumn string `bun:"-" json:"resource"`
}

type PathPermissions []*PathPermission

func (items PathPermissions) GetPathPermissionsRolesForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PathPermissionsRolesForDelete...)
	}
	return itemsForGet
}

func (items PathPermissions) GetPathPermissionsRoles() PathPermissionsRoles {
	itemsForGet := make(PathPermissionsRoles, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PathPermissionsRoles...)
	}
	return itemsForGet
}
