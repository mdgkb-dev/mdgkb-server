package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Regalia struct {
	bun.BaseModel `bun:"regalias,alias:regalias"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`

	DoctorID uuid.NullUUID `bun:"type:uuid" json:"doctorId"`
	Doctor   *Doctor       `bun:"rel:belongs-to" json:"doctor"`

	HeadID uuid.NullUUID `bun:"type:uuid" json:"headId"`
	Head   *Head         `bun:"rel:belongs-to" json:"head"`
}

type Regalias []*Regalia
