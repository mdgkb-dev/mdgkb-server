package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Form struct {
	bun.BaseModel `bun:"forms,alias:forms"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" json:"id,omitempty"`

	Fields          Fields      `bun:"rel:has-many" json:"fields"`
	FieldsForDelete []uuid.UUID `bun:"-" json:"fieldsForDelete"`
}

type Forms []*Form

func (item *Form) SetIdForChildren() {
	for i := range item.Fields {
		item.Fields[i].FormID = item.ID
	}
}

func (items Forms) SetIdForChildren() {
	for i := range items {
		items[i].SetIdForChildren()
	}
}
