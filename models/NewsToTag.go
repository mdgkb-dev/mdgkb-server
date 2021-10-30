package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type NewsToTag struct {
	bun.BaseModel `bun:"news_to_tags,alias:news_to_tags"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	NewsID        uuid.UUID `bun:"type:uuid" json:"newsId" `
	News          *News     `bun:"rel:belongs-to" json:"news" `
	TagID         uuid.UUID `bun:"type:uuid" json:"tagId" `
	Tag           *Tag      `bun:"rel:belongs-to" json:"tag" `
}

type NewsToTags []*NewsToTag
