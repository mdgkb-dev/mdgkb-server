package models

import "github.com/google/uuid"

type NewsToCategory struct {
	ID         uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	NewsID     uuid.NullUUID `bun:"type:uuid" json:"newsId"`
	News       *News         `bun:"rel:belongs-to"`
	CategoryID uuid.UUID     `bun:"type:uuid" json:"categoryId"`
	Category   *Category     `bun:"rel:belongs-to"`
}

type NewsToCategories []*NewsToCategory
