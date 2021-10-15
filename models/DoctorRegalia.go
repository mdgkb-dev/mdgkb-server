package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DoctorRegalia struct {
	bun.BaseModel `bun:"doctor_regalias,alias:doctor_regalias"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`

	DoctorID uuid.UUID `bun:"type:uuid" json:"doctorId"`
	Doctor   *Doctor   `bun:"rel:belongs-to" json:"doctor"`
}

type DoctorRegalias []*DoctorRegalia
