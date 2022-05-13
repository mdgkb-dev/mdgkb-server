package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Gate struct {
	bun.BaseModel    `bun:"gates,alias:gates"`
	ID               uuid.NullUUID    `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name             string           `json:"name"`
	ApplicationsCars ApplicationsCars `bun:"rel:has-many" json:"applicationsCars"`
}

type Gates []*Gate
