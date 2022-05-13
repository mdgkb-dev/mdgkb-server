package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PartnerType struct {
	bun.BaseModel `bun:"partner_types,alias:partner_types"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Slug          string        `json:"slug"`
	ShowImage     bool          `json:"showImage"`
}

type PartnerTypes []*PartnerType
