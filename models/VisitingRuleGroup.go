package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type VisitingRuleGroup struct {
	bun.BaseModel `bun:"visiting_rules_groups,alias:visiting_rules_groups"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `bun:"name" json:"name"`
	Color         string        `json:"color"`
	Order         uint          `bun:"visiting_rule_group_order" json:"order"`
	Division      *Division     `bun:"rel:belongs-to" json:"division"`
	DivisionID    uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"divisionId"`

	VisitingRules          VisitingRules `bun:"rel:has-many" json:"visitingRules"`
	VisitingRulesForDelete []uuid.UUID   `bun:"-" json:"visitingRulesForDelete"`
}

type VisitingRulesGroups []*VisitingRuleGroup

func (items VisitingRulesGroups) GetVisitingRules() VisitingRules {
	itemsForGet := make(VisitingRules, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.VisitingRules...)
	}
	return itemsForGet
}

func (items VisitingRulesGroups) GetVisitingRulesForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.VisitingRulesForDelete...)
	}
	return itemsForGet
}

func (item *VisitingRuleGroup) SetIDForChildren() {
	for i := range item.VisitingRules {
		item.VisitingRules[i].VisitingRuleGroupID = item.ID
	}
}

func (items VisitingRulesGroups) SetIDForChildren() {
	for i := range items {
		items[i].SetIDForChildren()
	}
}
