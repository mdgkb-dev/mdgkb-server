package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Page struct {
	bun.BaseModel `bun:"pages,alias:pages"`
	ID                  uuid.UUID      `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Title               string         `json:"title"`
	Content             string         `json:"content"`
	Slug                string         `json:"slug"`
	Link                string         `json:"link"`
}

type Pages []*Page