package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PaidProgramService struct {
	bun.BaseModel              `bun:"paid_program_services,alias:paid_program_services"`
	ID                         uuid.NullUUID             `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                       string                    `json:"name"`
	Quantity                   string                    `json:"quantity"`
	Order                      int                       `bun:"service_order" json:"order"`
	PaidProgramServicesGroup   *PaidProgramServicesGroup `bun:"rel:belongs-to" json:"paidProgramServicesGroup"`
	PaidProgramServicesGroupID uuid.NullUUID             `bun:"type:uuid" json:"paidProgramServicesGroupId"`
}

type PaidProgramServices []*PaidProgramService

func (item *PaidProgramService) SetIdForChildren() {
}

func (items PaidProgramServices) SetIdForChildren() {
	for i := range items {
		items[i].SetIdForChildren()
	}
}
