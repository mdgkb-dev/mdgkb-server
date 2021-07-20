package models

import "github.com/google/uuid"

type NewsComment struct {
	ID     uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	NewsID uuid.UUID `bun:"type:uuid" json:"newsId"`
	UserId uuid.UUID `bun:"type:uuid" json:"userId"`
	Text   string    `json:"text"`
}
