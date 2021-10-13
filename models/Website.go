package models

import (
	"github.com/google/uuid"
)

type Website struct {
	ID            uuid.UUID    `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	Address       string       `json:"address"`
	Description   string       `json:"description"`
	ContactInfo   *ContactInfo `bun:"rel:belongs-to" json:"contactInfo"`
	ContactInfoId uuid.UUID    `bun:"type:uuid"`
}

type Websites []*Website
