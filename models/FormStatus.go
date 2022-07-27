package models

import (
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"github.com/uptrace/bun"
)

type FormStatus struct {
	bun.BaseModel  `bun:"form_statuses,alias:form_statuses"`
	ID             uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name           string        `json:"name"`
	Label          string        `json:"label"`
	Color          string        `json:"color"`
	ModActionName  string        `json:"modActionName"`
	UserActionName string        `json:"userActionName"`
	IsEditable     bool          `json:"isEditable"`
	SendEmail      bool          `json:"sendEmail"`

	FormStatusToFormStatuses          FormStatusToFormStatuses `bun:"rel:has-many" json:"formStatusToFormStatuses"`
	FormStatusToFormStatusesForDelete []string                 `bun:"-" json:"formStatusToFormStatusesForDelete"`

	Icon   *FileInfo     `bun:"rel:belongs-to" json:"icon"`
	IconID uuid.NullUUID `bun:"type:uuid"  json:"iconId"`

	FormStatusGroup   *FormStatusGroup `bun:"rel:belongs-to" json:"formStatusGroup"`
	FormStatusGroupID uuid.NullUUID    `bun:"type:uuid"  json:"formStatusGroupId"`
}

type FormStatuses []*FormStatus

func (item *FormStatus) SetIDForChildren() {
	for i := range item.FormStatusToFormStatuses {
		item.FormStatusToFormStatuses[i].FormStatusID = item.ID
	}
}

func (items FormStatuses) SetIDForChildren() {
	for i := range items {
		items[i].SetIDForChildren()
	}
}

func (item *FormStatus) SetForeignKeys() {
	item.IconID.UUID = item.Icon.ID.UUID
	item.IconID = item.Icon.ID
	item.FormStatusGroupID = item.FormStatusGroup.ID
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

func (items FormStatuses) SetFilePath(fileID *string) *string {
	for i := range items {
		path := items[i].SetFilePath(fileID)
		if path != nil {
			return path
		}
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
