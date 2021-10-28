package models

import "github.com/google/uuid"

type Tag struct {
	ID    uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Color string    `json:"color"`
	Label string    `json:"label"`
}

type Tags []*Tag
