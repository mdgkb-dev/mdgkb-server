package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PreparationToTag struct {
	bun.BaseModel    `bun:"preparations_to_tags,alias:preparations_to_tags"`
	ID               uuid.NullUUID   `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Preparation      *Preparation    `bun:"rel:belongs-to" json:"preparation"`
	PreparationID    uuid.NullUUID   `bun:"type:uuid" json:"preparationId"`
	PreparationTag   *PreparationTag `bun:"rel:belongs-to" json:"preparationTag"`
	PreparationTagID uuid.NullUUID   `bun:"type:uuid" json:"preparationTagId"`
}

type PreparationsToTags []*PreparationToTag
