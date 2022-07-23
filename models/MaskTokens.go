package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MaskToken struct {
	bun.BaseModel `bun:"mask_tokes,alias:mask_tokes"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`
	Pattern       string    `json:"pattern"`
	Uppercase     bool      `json:"uppercase"`

	Field   *Field    `bun:"rel:belongs-to" json:"field"`
	FieldID uuid.UUID `bun:"type:uuid" json:"fieldId"`
}

type MaskTokens []*MaskToken
