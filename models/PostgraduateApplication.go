package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PostgraduateApplication struct {
	bun.BaseModel `bun:"postgraduate_applications,select:postgraduate_applications,alias:postgraduate_applications"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	PostgraduateCourse   *PostgraduateCourse `bun:"rel:belongs-to" json:"postgraduateCourse"`
	PostgraduateCourseID uuid.NullUUID       `bun:"type:uuid,nullzero,default:NULL" json:"postgraduateCourseId"`

	FormValue   *FormValue    `bun:"rel:belongs-to" json:"formValue"`
	FormValueID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"formValueId"`
}

type PostgraduateApplications []*PostgraduateApplication

func (item *PostgraduateApplication) SetForeignKeys() {
	item.PostgraduateCourseID = item.PostgraduateCourse.ID
	item.FormValueID = item.FormValue.ID
}

func (item *PostgraduateApplication) SetFilePath(fileID *string) *string {
	path := item.FormValue.SetFilePath(fileID)
	return path
}

func (item *PostgraduateApplication) SetIdForChildren() {
}
