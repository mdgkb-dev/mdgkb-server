package models

import (
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	ID          uuid.UUID    `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	UserId      uuid.UUID    `bun:"type:uuid" json:"userId"`
	Text        string       `json:"text"`
	PublishedOn time.Time    `json:"publishedOn"`
	User        *User        `bun:"rel:belongs-to" json:"user"`
}
