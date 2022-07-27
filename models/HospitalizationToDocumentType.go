package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type HospitalizationToDocumentType struct {
	bun.BaseModel `bun:"hospitalizations_to_document_types,alias:vacancies_responses_to_document_types"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id,omitempty"`

	DocumentType   *DocumentType `bun:"rel:belongs-to" json:"documentType"`
	DocumentTypeID uuid.UUID     `bun:"type:uuid"  json:"documentTypeId"`

	Hospitalization   *Hospitalization `bun:"rel:belongs-to" json:"hospitalization"`
	HospitalizationID uuid.UUID        `bun:"type:uuid"  json:"hospitalizationId"`
}

type HospitalizationsToDocumentTypes []*HospitalizationToDocumentType
