package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PreparationRule struct {
	bun.BaseModel           `bun:"preparations_rules,alias:preparations_rules"`
	ID                      uuid.UUID              `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                    string                 `json:"name"`
	Time                    string                 `bun:"rule_time" json:"time"`
	PreparationRulesGroup   *PreparationRulesGroup `bun:"rel:belongs-to" json:"preparationRulesGroup"`
	PreparationRulesGroupID uuid.NullUUID          `bun:"type:uuid" json:"preparationRulesGroupId"`
}

type PreparationRules []*PreparationRule
