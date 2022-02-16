package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Center struct {
	bun.BaseModel `bun:"centers,alias:centers"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Info          string        `json:"info"`
	Address       string        `json:"address"`
	Slug          string        `json:"slug"`
}

type Centers []*Center
