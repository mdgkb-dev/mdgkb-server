package models

import "github.com/google/uuid"

type NewsLike struct {
	ID     uuid.UUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	NewsID uuid.NullUUID `bun:"type:uuid" json:"newsId"`
	UserID uuid.UUID     `bun:"type:uuid" json:"userId"`
}

type NewsLikes []*NewsLike
