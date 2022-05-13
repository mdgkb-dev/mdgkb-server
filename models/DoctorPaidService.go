package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DoctorPaidService struct {
	bun.BaseModel `bun:"doctor_paid_services,alias:doctor_paid_services"`
	ID            uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	DoctorID      uuid.NullUUID `bun:"type:uuid" json:"doctorId"`
	Doctor        *Doctor       `bun:"rel:belongs-to" json:"doctor"`
	PaidService   *PaidService  `bun:"rel:belongs-to" json:"paidService"`
	PaidServiceID uuid.UUID     `bun:"type:uuid" json:"paidServiceId"`
}

type DoctorPaidServices []*DoctorPaidService
