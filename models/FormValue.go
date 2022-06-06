package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type FormValue struct {
	bun.BaseModel `bun:"form_values,alias:form_values"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	CreatedAt     time.Time     `json:"createdAt"`
	IsNew         bool          `json:"isNew"`
	EmailNotify   bool          `bun:"-" json:"emailNotify"`
	User          *User         `bun:"rel:belongs-to" json:"user"`
	UserID        uuid.UUID     `bun:"type:uuid" json:"userId"`

	Fields      Fields      `bun:"rel:has-many" json:"fields"`
	FieldValues FieldValues `bun:"rel:has-many" json:"fieldValues"`

	FormStatus   *FormStatus   `bun:"rel:belongs-to" json:"formStatus"`
	FormStatusID uuid.NullUUID `bun:"type:uuid" json:"formStatusId"`

	Child   *Child        `bun:"rel:belongs-to" json:"child"`
	ChildID uuid.NullUUID `bun:"type:uuid" json:"childId"`

	DpoApplication          *DpoApplication          `bun:"rel:has-one" json:"dpoApplication"`
	PostgraduateApplication *PostgraduateApplication `bun:"rel:has-one" json:"postgraduateApplication"`
	CandidateApplication    *CandidateApplication    `bun:"rel:has-one" json:"candidateApplication"`
	ResidencyApplication    *ResidencyApplication    `bun:"rel:has-one" json:"residencyApplication"`
}

type FormValues []*FormValue

func (item *FormValue) SetForeignKeys() {
	item.UserID = item.User.ID.UUID
	item.FormStatusID = item.FormStatus.ID
	if item.Child != nil {
		item.ChildID = item.Child.ID
	}
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
		filePath = item.FieldValues[i].FieldValuesFiles.SetFilePath(fileID)
		if filePath != nil {
			return filePath
		}
	}

	return nil
}

func (item *FormValue) GetFiles() []FileInfo {
	files := make([]FileInfo, 0)
	for i := range item.FieldValues {
		if item.FieldValues[i].File != nil && item.FieldValues[i].File.FileSystemPath != "" {
			files = append(files, *item.FieldValues[i].File)
		}
	}
	return files
}
