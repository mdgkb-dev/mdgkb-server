package models

import (
	"github.com/google/uuid"
)

type TelephoneNumber struct {
	ID            uuid.UUID    `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	Number        string       `json:"number"`
	Description   string       `json:"description"`
	ContactInfo   *ContactInfo `bun:"rel:belongs-to" json:"contactInfo"`
	ContactInfoId uuid.UUID    `bun:"type:uuid"`
}
