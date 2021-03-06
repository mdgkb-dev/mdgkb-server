package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Hospitalization struct {
	bun.BaseModel                   `bun:"hospitalizations,alias:hospitalizations"`
	ID                              uuid.UUID                       `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name                            string                          `json:"name"`
	Paid                            bool                            `json:"paid"`
	HospitalizationsToDocumentTypes HospitalizationsToDocumentTypes `bun:"rel:has-many" json:"hospitalizationsToDocumentTypes"`
}

type Hospitalizations []*Hospitalization
