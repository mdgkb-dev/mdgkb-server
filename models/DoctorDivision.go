package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DoctorDivision struct {
	bun.BaseModel `bun:"doctors_divisions,alias:doctors_divisions"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Doctor        *Doctor       `bun:"rel:belongs-to" json:"doctor"`
	DoctorID      uuid.NullUUID `bun:"type:uuid" json:"doctorId"`
	Division      *Division     `bun:"rel:belongs-to" json:"division"`
	DivisionID    uuid.NullUUID `bun:"type:uuid" json:"divisionId"`
	Show          bool          ` json:"show"`
}

type DoctorsDivisions []*DoctorDivision
