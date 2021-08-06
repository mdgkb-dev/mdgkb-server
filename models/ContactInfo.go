package models

import (
	"github.com/google/uuid"
)

type ContactInfo struct {
	ID               uuid.UUID          `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	Emails           []*Email           `bun:"rel:has-many" json:"emails"`
	TelephoneNumbers []*TelephoneNumber `bun:"rel:has-many" json:"telephoneNumbers"`
	Websites         []*Website         `bun:"rel:has-many" json:"websites"`
}
