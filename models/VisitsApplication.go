package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type VisitsApplication struct {
	bun.BaseModel `bun:"visits_applications,select:visits_applications_view,alias:visits_applications_view"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	WithCar       bool          `json:"withCar"`

	Division   *Division     `bun:"rel:belongs-to" json:"division"`
	DivisionID uuid.NullUUID `bun:"type:uuid" json:"divisionId,omitempty"`

	Gate   *Gate         `bun:"rel:belongs-to" json:"gate"`
	GateID uuid.NullUUID `bun:"type:uuid" json:"gateId,omitempty"`

	FormValue   *FormValue    `bun:"rel:belongs-to" json:"formValue"`
	FormValueID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"formValueId"`

	Visits          Visits      `bun:"rel:has-many" json:"visits"`
	VisitsForDelete []uuid.UUID `bun:"-" json:"visitsForDelete"`
}

type VisitsApplications []*VisitsApplication

func (item *VisitsApplication) SetForeignKeys() {
	item.DivisionID = item.Division.ID
	item.GateID = item.Gate.ID
	item.FormValueID = item.FormValue.ID
}

func (item *VisitsApplication) SetFilePath(fileID *string) *string {
	path := item.FormValue.SetFilePath(fileID)
	return path
}

func (item *VisitsApplication) SetIDForChildren() {
	for i := range item.Visits {
		item.Visits[i].VisitsApplicationID = item.ID
	}
}
