package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Position struct {
	bun.BaseModel `bun:"positions,alias:positions"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Order         int           `bun:"item_order" json:"order"`
	Show          bool          `json:"show"`
}

type Positions []*Position
