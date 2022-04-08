package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type FormStatus struct {
	bun.BaseModel  `bun:"form_statuses,alias:form_statuses"`
	ID             uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name           string        `json:"name"`
	Label          string        `json:"label"`
	Color          string        `json:"color"`
	ModActionName  string        `json:"modActionName"`
	UserActionName string        `json:"userActionName"`
	IsEditable     bool          `json:"isEditable"`

	FormStatusToFormStatuses FormStatusToFormStatuses `bun:"rel:has-many" json:"formStatusToFormStatuses"`
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

func (items FormStatuses) GetFormStatusToFormStatuses() FormStatusToFormStatuses {
	itemsForGet := make(FormStatusToFormStatuses, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].FormStatusToFormStatuses...)
	}
	return itemsForGet
}
