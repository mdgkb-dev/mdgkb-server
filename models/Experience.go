package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Experience struct {
	bun.BaseModel `bun:"experiences,alias:experiences"`
	ID            uuid.UUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	DoctorID      uuid.NullUUID `bun:"type:uuid" json:"doctorId"`
	Doctor        *Doctor       `bun:"rel:belongs-to" json:"doctor"`
	Start         int           `bun:"experience_start" json:"start"`
	End           int           `bun:"experience_end" json:"end"`
	Place         string        `json:"place"`
	Position      string        `json:"position"`
}

type Experiences []*Experience
