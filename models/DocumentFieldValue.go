package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DocumentFieldValue struct {
	bun.BaseModel `bun:"document_field_values,alias:document_field_values"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	ValueString   string    `json:"valueString"`
	ValueNumber   int       `json:"valueNumber"`
	ValueDate     time.Time `json:"valueDate"`

	Document   *Document `bun:"rel:has-one" json:"document"`
	DocumentID uuid.UUID `bun:"type:uuid" json:"documentId"`

	DocumentTypeField   *DocumentTypeField `bun:"rel:belongs-to" json:"documentTypeField"`
	DocumentTypeFieldID uuid.UUID          `bun:"type:uuid" json:"documentTypeFieldId"`
}

type DocumentFieldValues []*DocumentFieldValue
