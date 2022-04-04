package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type FormValue struct {
	bun.BaseModel `bun:"form_values,alias:form_values"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	CreatedAt     time.Time     `json:"createdAt"`
	IsNew         bool          `json:"isNew"`

	User   *User     `bun:"rel:belongs-to" json:"user"`
	UserID uuid.UUID `bun:"type:uuid" json:"userId"`

	Fields      Fields      `bun:"rel:has-many" json:"fields"`
	FieldValues FieldValues `bun:"rel:has-many" json:"fieldValues"`
	// FieldValuesForDelete []uuid.UUID `bun:"-" json:"fieldValuesForDelete"`
}

type FormValues []*FormValue

func (item *FormValue) SetForeignKeys() {
	item.UserID = item.User.ID
}

func (item *FormValue) SetIdForChildren() {
	for i := range item.Fields {
		item.Fields[i].FormValueID = item.ID
	}
	for i := range item.FieldValues {
		item.FieldValues[i].FormValueID = item.ID
		item.FieldValues[i].Field.FormValueID = item.ID
	}
}

func (item *FormValue) SetFilePath(fileID *string) *string {
	for i := range item.FieldValues {
		filePath := item.FieldValues[i].SetFilePath(fileID)
		if filePath != nil {
			return filePath
		}
	}
	return nil
}
