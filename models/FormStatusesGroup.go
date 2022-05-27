package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type FormStatusGroup struct {
	bun.BaseModel `bun:"form_status_groups,alias:form_status_groups"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	FormStatuses  FormStatuses  `bun:"rel:has-many" json:"formStatuses"`
}

type FormStatusGroups []*FormStatusGroup

func (item *FormStatusGroup) SetIdForChildren() {
	for i := range item.FormStatuses {
		item.FormStatuses[i].FormStatusGroupID = item.ID
	}
}

func (items FormStatusGroups) SetIdForChildren() {
	for i := range items {
		items[i].SetIdForChildren()
	}
}

func (items FormStatusGroups) GetFormStatuses() FormStatuses {
	itemsForGet := make(FormStatuses, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].FormStatuses...)
	}
	return itemsForGet
}

func (item *FormStatusGroup) SetFilePath(fileID *string) *string {
	return item.FormStatuses.SetFilePath(fileID)
}
