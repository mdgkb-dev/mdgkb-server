package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PreparationRulesGroup struct {
	bun.BaseModel `bun:"preparations_rules_groups,alias:preparations_rules_groups"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Order         int           `bun:"item_order" json:"order"`
	Preparation   *Preparation  `bun:"rel:belongs-to" json:"preparation"`
	PreparationID uuid.NullUUID `bun:"type:uuid" json:"preparationId"`

	PreparationRules          PreparationRules `bun:"rel:has-many" json:"preparationRules"`
	PreparationRulesForDelete []uuid.UUID      `bun:"-" json:"preparationRulesForDelete"`
}

type PreparationRulesGroups []*PreparationRulesGroup

func (items PreparationRulesGroups) GetPreparationRules() PreparationRules {
	itemsForGet := make(PreparationRules, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PreparationRules...)
	}
	return itemsForGet
}

func (items PreparationRulesGroups) GetPreparationRulesForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PreparationRulesForDelete...)
	}
	return itemsForGet
}

func (item *PreparationRulesGroup) SetIDForChildren() {
	for i := range item.PreparationRules {
		item.PreparationRules[i].PreparationRulesGroupID = item.ID
	}
}

func (items PreparationRulesGroups) SetIDForChildren() {
	for i := range items {
		items[i].SetIDForChildren()
	}
}
