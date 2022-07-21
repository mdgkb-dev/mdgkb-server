package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Role struct {
	bun.BaseModel `bun:"roles,alias:roles"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          RoleName      `json:"name"`
	Label         string        `json:"label"`
	StartPage     string        `json:"startPage"`
}

type Roles []*Role

type RoleName string

const (
	RoleNameUser  RoleName = "USER"
	RoleNameAdmin RoleName = "ADMIN"
)
