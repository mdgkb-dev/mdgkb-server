package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PreparationTag struct {
	bun.BaseModel `bun:"preparations_tags,alias:preparations_tags"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
}

type PreparationsTags []*PreparationTag
