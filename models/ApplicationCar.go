package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ApplicationCar struct {
	bun.BaseModel `bun:"applications_cars,alias:applications_cars"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	Division   *Division     `bun:"rel:belongs-to" json:"division"`
	DivisionID uuid.NullUUID `bun:"type:uuid" json:"divisionId,omitempty"`

	Gate   *Gate         `bun:"rel:belongs-to" json:"gate"`
	GateID uuid.NullUUID `bun:"type:uuid" json:"gateId,omitempty"`

	FormValue   *FormValue    `bun:"rel:belongs-to" json:"formValue"`
	FormValueID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"formValueId"`
}

type ApplicationsCars []*ApplicationCar

func (item *ApplicationCar) SetForeignKeys() {
	item.DivisionID = item.Division.ID
	item.GateID = item.Gate.ID
	item.FormValueID = item.FormValue.ID
}

func (item *ApplicationCar) SetFilePath(fileID *string) *string {
	path := item.FormValue.SetFilePath(fileID)
	return path
}
