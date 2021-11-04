package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DocumentField struct {
	bun.BaseModel `bun:"document_fields,alias:document_fields"`
	ID            uuid.UUID  `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string     `json:"name"`
	Order         uint       `bun:"document_field_order" json:"order"`
	DocumentID    uuid.UUID  `bun:"type:uuid" json:"documentTypeId"`
	ValueType     *ValueType `bun:"rel:belongs-to" json:"valueType"`
	ValueTypeID   uuid.UUID  `bun:"type:uuid" json:"valueTypeId"`
}

type DocumentFields []*DocumentField
