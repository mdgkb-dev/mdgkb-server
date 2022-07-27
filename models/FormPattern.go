package models

import (
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"github.com/uptrace/bun"
)

type FormPattern struct {
	bun.BaseModel             `bun:"form_patterns,alias:form_patterns"`
	Title                     string           `json:"title"`
	Code                      string           `json:"code"`
	FormStatusGroup           *FormStatusGroup `bun:"rel:belongs-to" json:"formStatusGroup"`
	FormStatusGroupID         uuid.NullUUID    `bun:"type:uuid" json:"formStatusGroupId"`
	DefaultFormStatus         *FormStatus      `bun:"rel:belongs-to" json:"defaultFormStatus"`
	DefaultFormStatusID       uuid.NullUUID    `bun:"type:uuid" json:"defaultFormStatusId"`
	WithPersonalDataAgreement bool             `json:"withPersonalDataAgreement"`
	PersonalDataAgreement     *FileInfo        `bun:"rel:belongs-to" json:"personalDataAgreement"`
	PersonalDataAgreementID   uuid.NullUUID    `bun:"type:uuid" json:"personalDataAgreementId"`
	Form
}

type FormPatterns []*FormPattern

func (item *FormPattern) SetForeignKeys() {
	if item.PersonalDataAgreement != nil {
		item.PersonalDataAgreementID = item.PersonalDataAgreement.ID
	}
	if item.FormStatusGroup != nil {
		item.FormStatusGroupID = item.FormStatusGroup.ID
	}
	if item.DefaultFormStatus != nil {
		item.DefaultFormStatusID = item.DefaultFormStatus.ID
	}
}

func (item *FormPattern) SetIDForChildren() {
	if len(item.Fields) == 0 {
		return
	}
	for i := range item.Fields {
		item.Fields[i].FormPatternID = item.ID
	}
}

func (item *FormPattern) SetFilePath(fileID *string) *string {
	if item.PersonalDataAgreement.ID.UUID.String() == *fileID {
		item.PersonalDataAgreement.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.PersonalDataAgreement.FileSystemPath
	}
	for i := range item.Fields {
		filePath := item.Fields[i].SetFilePath(fileID)
		if filePath != nil {
			return filePath
		}
	}
	return nil
}
