package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PaidProgramPackage struct {
	bun.BaseModel `bun:"paid_program_packages,alias:paid_program_packages"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Price         int           `json:"price"`
	Name          string        `json:"name"`
	PaidProgram   *PaidProgram  `bun:"rel:belongs-to" json:"paidProgram"`
	PaidProgramID uuid.NullUUID `bun:"type:uuid" json:"paidProgramId"`

	PaidProgramPackagesOptions          PaidProgramPackagesOptions `bun:"rel:has-many" json:"paidProgramPackagesOptions"`
	PaidProgramPackagesOptionsForDelete []uuid.UUID                `bun:"-" json:"paidProgramPackagesOptionsForDelete"`

	PaidProgramServicesGroups          PaidProgramServicesGroups `bun:"rel:has-many" json:"paidProgramServicesGroups"`
	PaidProgramServicesGroupsForDelete []uuid.UUID               `bun:"-" json:"paidProgramServicesGroupsForDelete"`
}

type PaidProgramPackages []*PaidProgramPackage

func (items PaidProgramPackages) SetIdForChildren() {
	for i := range items {
		items[i].SetIdForChildren()
	}
}

func (item *PaidProgramPackage) SetIdForChildren() {
	for i := range item.PaidProgramPackagesOptions {
		item.PaidProgramPackagesOptions[i].PaidProgramPackageID = item.ID
	}
	for i := range item.PaidProgramServicesGroups {
		item.PaidProgramServicesGroups[i].PaidProgramPackageID = item.ID
	}
}

func (items PaidProgramPackages) GetPaidProgramPackagesOptions() PaidProgramPackagesOptions {
	itemsForGet := make(PaidProgramPackagesOptions, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PaidProgramPackagesOptions...)
	}
	return itemsForGet
}

func (items PaidProgramPackages) GetPaidProgramPackagesOptionsForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PaidProgramPackagesOptionsForDelete...)
	}
	return itemsForGet
}

func (items PaidProgramPackages) GetPaidProgramServicesGroups() PaidProgramServicesGroups {
	itemsForGet := make(PaidProgramServicesGroups, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PaidProgramServicesGroups...)
	}
	return itemsForGet
}

func (items PaidProgramPackages) GetPaidProgramServicesGroupsForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PaidProgramServicesGroupsForDelete...)
	}
	return itemsForGet
}
