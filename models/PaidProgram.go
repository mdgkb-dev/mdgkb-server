package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PaidProgram struct {
	bun.BaseModel                     `bun:"paid_programs,alias:paid_programs"`
	ID                                uuid.NullUUID            `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                              string                   `json:"name"`
	Description                       string                   `json:"description"`
	Order                             int                      `bun:"program_order" json:""`
	PaidProgramOptionsGroups          PaidProgramOptionsGroups `bun:"rel:has-many" json:"paidProgramOptionsGroups"`
	PaidProgramOptionsGroupsForDelete []uuid.UUID              `bun:"-" json:"paidProgramOptionsGroupsForDelete"`

	PaidProgramsGroup   *PaidProgramsGroup `bun:"rel:belongs-to" json:"paidProgramsGroup"`
	PaidProgramsGroupID uuid.NullUUID      `bun:"type:uuid" json:"paidProgramsGroupId"`

	PaidProgramPackages          PaidProgramPackages `bun:"rel:has-many" json:"paidProgramPackages"`
	PaidProgramPackagesForDelete []uuid.UUID         `bun:"-" json:"paidProgramPackagesForDelete"`
}

type PaidPrograms []*PaidProgram

func (items PaidPrograms) SetIDForChildren() {
	for i := range items {
		items[i].SetIDForChildren()
	}
}

func (item *PaidProgram) SetIDForChildren() {
	for i := range item.PaidProgramPackages {
		item.PaidProgramPackages[i].PaidProgramID = item.ID
	}
	for i := range item.PaidProgramOptionsGroups {
		item.PaidProgramOptionsGroups[i].PaidProgramID = item.ID
	}
}

func (items PaidPrograms) GetPaidProgramPackages() PaidProgramPackages {
	itemsForGet := make(PaidProgramPackages, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PaidProgramPackages...)
	}
	return itemsForGet
}

func (items PaidPrograms) GetPaidProgramPackagesForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PaidProgramPackagesForDelete...)
	}
	return itemsForGet
}

func (items PaidPrograms) GetPaidProgramOptionsGroups() PaidProgramOptionsGroups {
	itemsForGet := make(PaidProgramOptionsGroups, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PaidProgramOptionsGroups...)
	}
	return itemsForGet
}

func (items PaidPrograms) GetPaidProgramOptionsGroupsForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PaidProgramOptionsGroupsForDelete...)
	}
	return itemsForGet
}
