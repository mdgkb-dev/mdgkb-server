package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ResidencyApplication struct {
	bun.BaseModel `bun:"residency_applications,select:residency_applications,alias:residency_applications"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	ResidencyCourse   *ResidencyCourse `bun:"rel:belongs-to" json:"residencyCourse"`
	ResidencyCourseID uuid.NullUUID    `bun:"type:uuid,nullzero,default:NULL" json:"residencyCourseId"`

	FormValue   *FormValue    `bun:"rel:belongs-to" json:"formValue"`
	FormValueID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"formValueId"`
}

type ResidencyApplications []*ResidencyApplication

func (item *ResidencyApplication) SetForeignKeys() {
	item.ResidencyCourseID = item.ResidencyCourse.ID
	item.FormValueID = item.FormValue.ID
}

func (item *ResidencyApplication) SetFilePath(fileID *string) *string {
	path := item.FormValue.SetFilePath(fileID)
	return path
}

func (item *ResidencyApplication) SetIdForChildren() {
}
