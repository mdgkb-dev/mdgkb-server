package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DonorRuleUser struct {
	bun.BaseModel `bun:"donor_rules_users,alias:donor_rules_users"`
	ID            uuid.UUID  `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	DonorRule     *DonorRule `bun:"rel:belongs-to" json:"donorRule"`
	DonorRuleID   uuid.UUID  `bun:"type:uuid" json:"donorRuleId"`
	User          *User      `bun:"rel:belongs-to" json:"user"`
	UserID        uuid.UUID  `bun:"type:uuid" json:"userId"`
}

type DonorRulesUsers []*DonorRuleUser
