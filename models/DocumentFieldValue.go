package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DocumentFieldValue struct {
	bun.BaseModel `bun:"document_field_value,alias:document_field_value"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	ValueString   string    `json:"valueString"`
	ValueNumber   int       `json:"valueNumber"`
	ValueDate     time.Time `json:"valueDate"`

	DocumentField   *DocumentField `bun:"rel:belongs-to" json:"documentField"`
	DocumentFieldID uuid.UUID      `bun:"type:uuid" json:"documentFieldId"`
}
