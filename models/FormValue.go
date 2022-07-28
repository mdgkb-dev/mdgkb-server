package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type FormValue struct {
	bun.BaseModel `bun:"form_values,alias:form_values"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	CreatedAt     time.Time     `json:"createdAt"`
	IsNew         bool          `json:"isNew"`
	ViewedByUser  bool          `json:"viewedByUser"`
	EmailNotify   bool          `bun:"-" json:"emailNotify"`
	ModComment    string        `json:"modComment"`
	User          *User         `bun:"rel:belongs-to" json:"user"`
	UserID        uuid.NullUUID `bun:"type:uuid" json:"userId"`
	ApprovingDate *time.Time    `json:"approvingDate"`
	Fields        Fields        `bun:"rel:has-many" json:"fields"`
	FieldValues   FieldValues   `bun:"rel:has-many" json:"fieldValues"`

	FormStatus   *FormStatus   `bun:"rel:belongs-to" json:"formStatus"`
	FormStatusID uuid.NullUUID `bun:"type:uuid" json:"formStatusId"`

	Child   *Child        `bun:"rel:belongs-to" json:"child"`
	ChildID uuid.NullUUID `bun:"type:uuid" json:"childId"`

	DpoApplication          *DpoApplication          `bun:"rel:has-one" json:"dpoApplication"`
	PostgraduateApplication *PostgraduateApplication `bun:"rel:has-one" json:"postgraduateApplication"`
	CandidateApplication    *CandidateApplication    `bun:"rel:has-one" json:"candidateApplication"`
	ResidencyApplication    *ResidencyApplication    `bun:"rel:has-one" json:"residencyApplication"`
	VisitsApplication       *VisitsApplication       `bun:"rel:has-one" json:"visitsApplication"`
	VacancyResponse         *VacancyResponse         `bun:"rel:has-one" json:"vacancyResponse"`
}

type FormValues []*FormValue

func (item *FormValue) SetForeignKeys() {
	item.UserID = item.User.ID
	item.FormStatusID = item.FormStatus.ID
	if item.Child != nil {
		item.ChildID = item.Child.ID
	}
}

func (item *FormValue) SetIDForChildren() {
	for i := range item.Fields {
		item.Fields[i].FormValueID = item.ID
	}
	for i := range item.FieldValues {
		item.FieldValues[i].FormValueID = item.ID
		if item.FieldValues[i].Field != nil {
			item.FieldValues[i].Field.FormValueID = item.ID
		}
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
	if item.ResidencyApplication != nil {
		filePath := item.ResidencyApplication.SetFilePath(fileID)
		if filePath != nil {
			return filePath
		}
	}
	return nil
}

func (item *FormValue) GetFiles() FileInfos {
	files := make(FileInfos, 0)
	for i := range item.FieldValues {
		if item.FieldValues[i].File != nil && item.FieldValues[i].File.FileSystemPath != "" {
			files = append(files, item.FieldValues[i].File)
		}
	}
	return files
}

func (item *FormValue) GetFieldValueByCode(code string) interface{} {
	var value interface{} = ""
	for _, fieldValue := range item.FieldValues {
		if fieldValue.Field.Code == code {
			if fieldValue.Field.ValueType.Name == "string" {
				value = fieldValue.ValueString
			}
			if fieldValue.Field.ValueType.Name == "number" {
				value = fieldValue.ValueNumber
			}
			if fieldValue.Field.ValueType.Name == "date" {
				value = fieldValue.ValueDate
			}
			break
		}
	}
	return value
}

func (item *FormValue) NormalizeDateFields() {
	for i := range item.FieldValues {
		if item.FieldValues[i].ValueDate != nil {
			d := item.FieldValues[i].ValueDate.Add(time.Hour * 3)
			item.FieldValues[i].ValueDate = &d
		}
	}
}
