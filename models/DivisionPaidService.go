package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DivisionPaidService struct {
	bun.BaseModel `bun:"division_paid_services,alias:division_paid_services"`
	ID            uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	DivisionID    uuid.NullUUID `bun:"type:uuid" json:"divisionId"`
	Division      *Division     `bun:"rel:belongs-to" json:"division"`
	PaidService   *PaidService  `bun:"rel:belongs-to" json:"paidService"`
	PaidServiceID uuid.UUID     `bun:"type:uuid" json:"paidServiceId"`
}

type DivisionPaidServices []*DivisionPaidService
