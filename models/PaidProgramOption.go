package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PaidProgramOption struct {
	bun.BaseModel             `bun:"paid_program_options,alias:paid_program_options"`
	ID                        uuid.UUID                `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                      string                   `json:"name"`
	Order                     uint                     `bun:"item_order" json:"order"`
	PaidProgramOptionsGroup   *PaidProgramOptionsGroup `bun:"rel:belongs-to" json:"paidProgramOptionsGroup"`
	PaidProgramOptionsGroupID uuid.NullUUID            `bun:"type:uuid" json:"paidProgramOptionsGroupId"`
}

type PaidProgramOptions []*PaidProgramOption
