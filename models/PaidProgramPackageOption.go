package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PaidProgramPackageOption struct {
	bun.BaseModel        `bun:"paid_programs_packages_options,alias:paid_programs_packages_options"`
	ID                   uuid.NullUUID       `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	PaidProgramOption    *PaidProgramOption  `bun:"rel:belongs-to" json:"paidProgramOption"`
	PaidProgramOptionID  uuid.UUID           `bun:"type:uuid" json:"paidProgramOptionId"`
	PaidProgramPackage   *PaidProgramPackage `bun:"rel:belongs-to" json:"paidProgramPackage"`
	PaidProgramPackageID uuid.NullUUID       `bun:"type:uuid" json:"paidProgramPackageId"`
}

type PaidProgramPackagesOptions []*PaidProgramPackageOption
