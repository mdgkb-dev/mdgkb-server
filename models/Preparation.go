package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Preparation struct {
	bun.BaseModel `bun:"preparations,alias:preparations"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	OMS           bool          `json:"oms"`
	DMS           bool          `json:"dms"`
	Laboratory    bool          `json:"laboratory"`

	PreparationRulesGroups          PreparationRulesGroups `bun:"rel:has-many" json:"preparationRulesGroups"`
	PreparationRulesGroupsForDelete []uuid.UUID            `bun:"-" json:"preparationRulesGroupsForDelete"`
}

type Preparations []*Preparation

type PreparationsWithCount struct {
	Preparations Preparations `json:"items"`
	Count        int          `json:"count"`
}

func (item *Preparation) SetIDForChildren() {
	for i := range item.PreparationRulesGroups {
		item.PreparationRulesGroups[i].PreparationID = item.ID
	}
}

func (items Preparations) SetIDForChildren() {
	for i := range items {
		items[i].SetIDForChildren()
	}
}

func (items Preparations) GetPreparationRulesGroups() PreparationRulesGroups {
	itemsForGet := make(PreparationRulesGroups, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PreparationRulesGroups...)
	}
	return itemsForGet
}

func (items Preparations) GetPreparationRulesGroupsForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PreparationRulesGroupsForDelete...)
	}
	return itemsForGet
}
