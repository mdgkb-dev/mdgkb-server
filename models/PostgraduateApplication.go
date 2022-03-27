package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PostgraduateApplication struct {
	bun.BaseModel `bun:"postgraduate_applications,select:postgraduate_applications,alias:postgraduate_applications"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	CreatedAt     time.Time     `json:"createdAt"`

	PostgraduateCourse   *PostgraduateCourse `bun:"rel:belongs-to" json:"postgraduateCourse"`
	PostgraduateCourseID uuid.NullUUID       `bun:"type:uuid,nullzero,default:NULL" json:"postgraduateCourseId"`

	User   *User     `bun:"rel:belongs-to" json:"user"`
	UserID uuid.UUID `bun:"type:uuid" json:"userId"`

	FieldValues FieldValues `bun:"rel:has-many" json:"fieldValues"`
	IsNew       bool        `json:"isNew"`
}

type PostgraduateApplications []*PostgraduateApplication

func (item *PostgraduateApplication) SetForeignKeys() {
	item.UserID = item.User.ID
	item.PostgraduateCourseID = item.PostgraduateCourse.ID
}

func (item *PostgraduateApplication) SetFilePath(fileID *string) *string {
	for i := range item.FieldValues {
		filePath := item.FieldValues[i].SetFilePath(fileID)
		if filePath != nil {
			return filePath
		}
	}
	return nil
}

func (item *PostgraduateApplication) SetIdForChildren() {
	for i := range item.FieldValues {
		item.FieldValues[i].PostgraduateApplicationID = item.ID
	}
}
