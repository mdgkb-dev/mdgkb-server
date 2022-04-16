package models

import (
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"github.com/uptrace/bun"
)

type FormStatus struct {
	bun.BaseModel                     `bun:"form_statuses,alias:form_statuses"`
	ID                                uuid.NullUUID            `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                              string                   `json:"name"`
	Label                             string                   `json:"label"`
	Color                             string                   `json:"color"`
	ModActionName                     string                   `json:"modActionName"`
	UserActionName                    string                   `json:"userActionName"`
	IsEditable                        bool                     `json:"isEditable"`
	Icon                              *FileInfo                `bun:"rel:belongs-to" json:"icon"`
	IconId                            uuid.NullUUID            `bun:"type:uuid"  json:"iconId"`
	SendEmail                         bool                     `json:"sendEmail"`
	FormStatusToFormStatuses          FormStatusToFormStatuses `bun:"rel:has-many" json:"formStatusToFormStatuses"`
	FormStatusToFormStatusesForDelete []string                 `bun:"-" json:"formStatusToFormStatusesForDelete"`
}

type FormStatuses []*FormStatus

func (item *FormStatus) SetIdForChildren() {
	for i := range item.FormStatusToFormStatuses {
		item.FormStatusToFormStatuses[i].FormStatusID = item.ID
	}
}

func (items FormStatuses) SetIdForChildren() {
	for i := range items {
		items[i].SetIdForChildren()
	}
}

func (item *FormStatus) SetForeignKeys() {
	item.IconId.UUID = item.Icon.ID.UUID
	item.IconId = item.Icon.ID
}

func (items FormStatuses) GetFormStatusToFormStatuses() FormStatusToFormStatuses {
	itemsForGet := make(FormStatusToFormStatuses, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].FormStatusToFormStatuses...)
	}
	return itemsForGet
}

func (item *FormStatus) SetFilePath(fileID *string) *string {
	if item.Icon.ID.UUID.String() == *fileID {
		item.Icon.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.Icon.FileSystemPath
	}
	return nil
}

func (items FormStatuses) GetFormStatusToFormStatusesForDelete() []string {
	itemsForGet := make([]string, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.FormStatusToFormStatusesForDelete...)
	}
	return itemsForGet
}
