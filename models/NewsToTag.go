package models

import "github.com/google/uuid"

type NewsToTag struct {
	ID     uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	NewsID uuid.UUID `bun:"type:uuid" json:"newsId" `
	News   *News     `bun:"rel:has-one" json:"news" `
	TagID  uuid.UUID `bun:"type:uuid" json:"tagId" `
	Tag    *Tag      `bun:"rel:has-one" json:"tag" `
}

type NewsToTags []*NewsToTag
