package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type VisitingRule struct {
	bun.BaseModel `bun:"visiting_rules,alias:visiting_rules"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Order         uint      `bun:"rule_order" json:"order"`
	Text          string    `json:"text"`
	IsListItem    bool      `json:"isListItem"`

	Division   *Division     `bun:"rel:belongs-to" json:"division"`
	DivisionID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"divisionId"`
}

type VisitingRules []*VisitingRule

type VisitingRulesWithDeleted struct {
	VisitingRules          VisitingRules `json:"visitingRules"`
	VisitingRulesForDelete []uuid.UUID   `json:"visitingRulesForDelete"`
}

func (item *VisitingRule) SetForeignKeys() {
	item.DivisionID = item.Division.ID
}
