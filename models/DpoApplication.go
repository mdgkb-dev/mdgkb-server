package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DpoApplication struct {
	bun.BaseModel `bun:"dpo_applications,select:dpo_applications_view,alias:dpo_applications_view"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	NmoCourse   *NmoCourse    `bun:"rel:belongs-to" json:"nmoCourse"`
	NmoCourseID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"nmoCourseId"`

	FormValue   *FormValue    `bun:"rel:belongs-to" json:"formValue"`
	FormValueID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"formValueId"`
}

type DpoApplications []*DpoApplication

type DpoApplicationsWithCount struct {
	DpoApplications DpoApplications `json:"dpoApplications"`
	Count           int             `json:"count"`
}

func (item *DpoApplication) SetForeignKeys() {
	item.NmoCourseID = item.NmoCourse.ID
	item.FormValueID = item.FormValue.ID
}

func (item *DpoApplication) SetFilePath(fileID *string) *string {
	path := item.FormValue.SetFilePath(fileID)
	return path
}

func (item *DpoApplication) SetIDForChildren() {
}
