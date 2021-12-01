package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Field struct {
	bun.BaseModel `bun:"fields,alias:fields"`
	ID            uuid.UUID  `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string     `json:"name"`
	Order         uint       `bun:"field_order" json:"order"`
	FormID        uuid.UUID  `bun:"type:uuid" json:"formId"`
	ValueType     *ValueType `bun:"rel:belongs-to" json:"valueType"`
	ValueTypeID   uuid.UUID  `bun:"type:uuid" json:"valueTypeId"`
}

type Fields []*Field
