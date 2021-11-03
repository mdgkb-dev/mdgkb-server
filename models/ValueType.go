package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ValueType struct {
	bun.BaseModel `bun:"value_types,alias:value_types"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`
	ValueRelation string    `bun:"type:value_type_value_relation_enum" json:"valueRelation"`
}

type ValueTypes []*ValueType
