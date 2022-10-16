package models

import (
	"github.com/google/uuid"
)

type Email struct {
	ID            uuid.UUID    `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Address       string       `json:"address"`
	Description   string       `json:"description"`
	ContactInfo   *ContactInfo `bun:"rel:belongs-to" json:"contactInfo"`
	ContactInfoID uuid.UUID    `bun:"type:uuid" json:"contactInfoId"`
	Main          bool         `json:"main"`
}

type Emails []*Email
