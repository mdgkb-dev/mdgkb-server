package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PaidProgramOptionsGroup struct {
	bun.BaseModel `bun:"paid_program_options_groups,alias:paid_program_options_groups"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Order         uint          `bun:"item_order" json:"order"`

	PaidProgram   *PaidProgram  `bun:"rel:belongs-to" json:"paidProgram"`
	PaidProgramID uuid.NullUUID `bun:"type:uuid" json:"paidProgramId"`

	PaidProgramOptions          PaidProgramOptions `bun:"rel:has-many" json:"paidProgramOptions"`
	PaidProgramOptionsForDelete []uuid.UUID        `bun:"-" json:"paidProgramOptionsForDelete"`
}

type PaidProgramOptionsGroups []*PaidProgramOptionsGroup

func (item *PaidProgramOptionsGroup) SetIdForChildren() {
	for i := range item.PaidProgramOptions {
		item.PaidProgramOptions[i].PaidProgramOptionsGroupID = item.ID
	}
}

func (items PaidProgramOptionsGroups) SetIdForChildren() {
	for i := range items {
		items[i].SetIdForChildren()
	}
}

func (items PaidProgramOptionsGroups) GetPaidProgramOptions() PaidProgramOptions {
	itemsForGet := make(PaidProgramOptions, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PaidProgramOptions...)
	}
	return itemsForGet
}

func (items PaidProgramOptionsGroups) PaidProgramOptionsForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PaidProgramOptionsForDelete...)
	}
	return itemsForGet
}
