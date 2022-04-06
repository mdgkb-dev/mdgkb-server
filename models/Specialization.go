package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Specialization struct {
	bun.BaseModel `bun:"specializations,alias:specializations"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Slug          string        `json:"slug"`
	Code          string        `json:"code"`
}

type Specializations []*Specialization
