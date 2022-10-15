package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type VisitingRule struct {
	bun.BaseModel `bun:"visiting_rules,alias:visiting_rules"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Order         uint          `bun:"rule_order" json:"order"`
	Text          string        `json:"text"`
	IsListItem    bool          `json:"isListItem"`

	VisitingRuleGroup   *VisitingRuleGroup `bun:"rel:belongs-to" json:"VisitingRuleGroup"`
	VisitingRuleGroupID uuid.NullUUID      `bun:"type:uuid,nullzero,default:NULL" json:"VisitingRuleGroupID"`
}

type VisitingRules []*VisitingRule

type VisitingRulesWithDeleted struct {
	VisitingRules          VisitingRules `json:"visitingRules"`
	VisitingRulesForDelete []uuid.UUID   `json:"visitingRulesForDelete"`
}

func (item *VisitingRule) SetForeignKeys() {
	item.VisitingRuleGroupID = item.VisitingRuleGroup.ID
}
