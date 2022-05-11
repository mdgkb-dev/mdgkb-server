package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Tag struct {
	bun.BaseModel `bun:"tags,alias:tags"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Color         string    `json:"color"`
	Label         string    `json:"label"`
}

type Tags []*Tag
