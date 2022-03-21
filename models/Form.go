package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Form struct {
	bun.BaseModel `bun:"forms,alias:forms"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id,omitempty"`

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

func (item *Form) SetFilePath(fileID *string) *string {
	for i := range item.Fields {
		filePath := item.Fields[i].SetFilePath(fileID)
		if filePath != nil {
			return filePath
		}
	}
	return nil
}
