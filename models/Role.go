package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Role struct {
	bun.BaseModel `bun:"roles,alias:roles"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          RoleName      `json:"name"`
}

type Roles []*Role

type RoleName string

const (
	RoleNameUser  RoleName = "user"
	RoleNameAdmin RoleName = "admin"
)
