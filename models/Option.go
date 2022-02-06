package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Option struct {
	bun.BaseModel `bun:"options,alias:doctors_view"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Value         string        `json:"value"`
	Label         string        `json:"label"`
}

type Options []*Option
