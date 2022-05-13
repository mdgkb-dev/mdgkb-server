package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PaidProgramsGroup struct {
	bun.BaseModel         `bun:"paid_programs_groups,alias:paid_programs_groups"`
	ID                    uuid.NullUUID  `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                  string         `json:"name"`
	Order                 int            `bun:"group_order" json:"order"`
	PaidPrograms          []*PaidProgram `bun:"rel:has-many" json:"paidPrograms"`
	PaidProgramsForDelete []uuid.UUID    `bun:"-" json:"paidProgramsForDelete"`

	PaidProgramServices          PaidProgramServices `bun:"rel:has-many" json:"paidProgramServices"`
	PaidProgramServicesForDelete []uuid.UUID         `bun:"-" json:"paidProgramServicesForDelete"`
}

type PaidProgramsGroups []*PaidProgramsGroup

type PaidProgramsGroupsStruct struct {
	PaidProgramsGroups          PaidProgramsGroups `json:"paidProgramsGroups"`
	PaidProgramsGroupsForDelete []uuid.UUID        `json:"paidProgramsGroupsForDelete"`
}

func (item *PaidProgramsGroup) SetIdForChildren() {
	for i := range item.PaidPrograms {
		item.PaidPrograms[i].PaidProgramsGroupID = item.ID
	}
}

func (items PaidProgramsGroups) SetIdForChildren() {
	for i := range items {
		items[i].SetIdForChildren()
	}
}

func (items PaidProgramsGroups) GetPaidPrograms() PaidPrograms {
	itemsForGet := make(PaidPrograms, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PaidPrograms...)
	}
	return itemsForGet
}

func (items PaidProgramsGroups) GetPaidProgramsForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PaidProgramsForDelete...)
	}
	return itemsForGet
}
