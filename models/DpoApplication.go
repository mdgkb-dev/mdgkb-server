package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DpoApplication struct {
	bun.BaseModel `bun:"dpo_applications,select:dpo_applications_view,alias:dpo_applications_view"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	DpoCourse   *DpoCourse    `bun:"rel:belongs-to" json:"dpoCourse"`
	DpoCourseID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"dpoCourseId"`

	FormValue   *FormValue    `bun:"rel:belongs-to" json:"formValue"`
	FormValueID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"formValueId"`
}

type DpoApplications []*DpoApplication

func (item *DpoApplication) SetForeignKeys() {
	item.DpoCourseID = item.DpoCourse.ID
	item.FormValueID = item.FormValue.ID
}

func (item *DpoApplication) SetFilePath(fileID *string) *string {
	path := item.FormValue.SetFilePath(fileID)
	return path
}

func (item *DpoApplication) SetIdForChildren() {
}
