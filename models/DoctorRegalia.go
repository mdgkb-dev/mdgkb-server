package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Regalia struct {
	bun.BaseModel `bun:"regalias,alias:regalias"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`

	DoctorID uuid.UUID `bun:"type:uuid" json:"doctorId"`
	Doctor   *Doctor   `bun:"rel:belongs-to" json:"doctor"`

	HeadID uuid.UUID `bun:"type:uuid" json:"headId"`
	Head   *Head   `bun:"rel:belongs-to" json:"head"`
}

type Regalias []*Regalia
