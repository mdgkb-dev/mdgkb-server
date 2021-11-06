package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DocumentFieldValueToHuman struct {
	bun.BaseModel   `bun:"document_field_value_to_humans,alias:document_field_value_to_humans"`
	ID              uuid.UUID          `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	DocumentField   *DocumentTypeField `bun:"rel:belongs-to" json:"documentField"`
	DocumentFieldID uuid.UUID          `bun:"type:uuid" json:"documentFieldId"`

	Human   *Human    `bun:"rel:belongs-to" json:"human"`
	HumanId uuid.UUID `bun:"type:uuid" json:"humanId"`
}
