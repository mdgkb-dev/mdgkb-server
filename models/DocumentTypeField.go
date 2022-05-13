package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DocumentTypeField struct {
	bun.BaseModel  `bun:"document_type_fields,alias:document_type_fields"`
	ID             uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name           string        `json:"name"`
	Order          uint          `bun:"document_type_field_order" json:"order"`
	DocumentTypeID uuid.NullUUID `bun:"type:uuid" json:"documentTypeId"`
	ValueType      *ValueType    `bun:"rel:belongs-to" json:"valueType"`
	ValueTypeID    uuid.UUID     `bun:"type:uuid" json:"valueTypeId"`
}

type DocumentTypeFields []*DocumentTypeField
