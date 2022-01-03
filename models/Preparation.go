package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Preparation struct {
	bun.BaseModel `bun:"preparations,alias:preparations"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`

	PreparationRulesGroups          PreparationRulesGroups `bun:"rel:has-many" json:"preparationRulesGroups"`
	PreparationRulesGroupsForDelete []uuid.UUID            `bun:"-" json:"preparationRulesGroupsForDelete"`

	PreparationsToTags          PreparationsToTags `bun:"rel:has-many" json:"preparationsToTags"`
	PreparationsToTagsForDelete []uuid.UUID        `bun:"-" json:"preparationsToTagsForDelete"`
}

type Preparations []*Preparation

func (item *Preparation) SetIdForChildren() {
	for i := range item.PreparationRulesGroups {
		item.PreparationRulesGroups[i].PreparationID = item.ID
	}
	for i := range item.PreparationsToTags {
		item.PreparationsToTags[i].PreparationID = item.ID
	}
}

func (items Preparations) SetIdForChildren() {
	for i := range items {
		items[i].SetIdForChildren()
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

func (items Preparations) GetPreparationsToTags() PreparationsToTags {
	itemsForGet := make(PreparationsToTags, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PreparationsToTags...)
	}
	return itemsForGet
}

func (items Preparations) GetPreparationsToTagsForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PreparationsToTagsForDelete...)
	}
	return itemsForGet
}
