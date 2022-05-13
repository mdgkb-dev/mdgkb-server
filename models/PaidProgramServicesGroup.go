package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PaidProgramServicesGroup struct {
	bun.BaseModel        `bun:"paid_program_services_groups,alias:paid_program_services_groups"`
	ID                   uuid.NullUUID       `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                 string              `json:"name"`
	Home                 bool                `json:"home"`
	Order                int                 `bun:"group_order" json:"order"`
	PaidProgramPackage   *PaidProgramPackage `bun:"rel:belongs-to" json:"paidProgramPackage"`
	PaidProgramPackageID uuid.NullUUID       `bun:"type:uuid" json:"paidProgramPackageId"`

	PaidProgramServices          PaidProgramServices `bun:"rel:has-many" json:"paidProgramServices"`
	PaidProgramServicesForDelete []uuid.UUID         `bun:"-" json:"paidProgramServicesForDelete"`
}

type PaidProgramServicesGroups []*PaidProgramServicesGroup

func (item *PaidProgramServicesGroup) SetIdForChildren() {
	for i := range item.PaidProgramServices {
		item.PaidProgramServices[i].PaidProgramServicesGroupID = item.ID
	}
}

func (items PaidProgramServicesGroups) SetIdForChildren() {
	for i := range items {
		items[i].SetIdForChildren()
	}
}

func (items PaidProgramServicesGroups) GetPaidProgramServices() PaidProgramServices {
	itemsForGet := make(PaidProgramServices, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PaidProgramServices...)
	}
	return itemsForGet
}

func (items PaidProgramServicesGroups) GetPaidProgramServicesForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PaidProgramServicesForDelete...)
	}
	return itemsForGet
}
