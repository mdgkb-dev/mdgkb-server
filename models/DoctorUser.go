package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DoctorUser struct {
	bun.BaseModel `bun:"doctors_users,alias:doctors_users"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Doctor        *Doctor   `bun:"rel:belongs-to" json:"doctor"`
	DoctorID      uuid.UUID `bun:"type:uuid" json:"doctorId"`
	User          *User     `bun:"rel:belongs-to" json:"user"`
	UserID        uuid.UUID `bun:"type:uuid" json:"userId"`
}

type DoctorsUsers []*DoctorUser
