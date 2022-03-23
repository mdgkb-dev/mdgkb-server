package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type DpoApplication struct {
	bun.BaseModel `bun:"dpo_applications,select:dpo_applications_view,alias:dpo_applications_view"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	CreatedAt     time.Time     `json:"createdAt"`

	DpoCourse   *DpoCourse    `bun:"rel:belongs-to" json:"dpoCourse"`
	DpoCourseID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"dpoCourseId"`

	User   *User     `bun:"rel:belongs-to" json:"user"`
	UserID uuid.UUID `bun:"type:uuid" json:"userId"`

	FieldValues FieldValues `bun:"rel:has-many" json:"fieldValues"`
	IsNew       bool        `json:"isNew"`
}

type DpoApplications []*DpoApplication

func (item *DpoApplication) SetForeignKeys() {
	item.UserID = item.User.ID
	item.DpoCourseID = item.DpoCourse.ID
}

func (item *DpoApplication) SetFilePath(fileID *string) *string {
	for i := range item.FieldValues {
		filePath := item.FieldValues[i].SetFilePath(fileID)
		if filePath != nil {
			return filePath
		}
	}
	return nil
}

func (item *DpoApplication) SetIdForChildren() {
	for i := range item.FieldValues {
		item.FieldValues[i].DpoApplicationID = item.ID
	}
}
