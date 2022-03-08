package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ApplicationCar struct {
	bun.BaseModel `bun:"applications_cars,alias:applications_cars"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	CarNumber     string        `json:"carNumber"`
	CarBrand      string        `json:"carBrand"`

	Division   *Division     `bun:"rel:belongs-to" json:"division"`
	DivisionID uuid.NullUUID `bun:"type:uuid" json:"divisionId,omitempty"`

	Gate   *Gate         `bun:"rel:belongs-to" json:"gate"`
	GateID uuid.NullUUID `bun:"type:uuid" json:"gateId,omitempty"`

	Date time.Time `bun:"hospitalization_date" json:"date"`

	User      *User     `bun:"rel:belongs-to" json:"user"`
	UserID    uuid.UUID `bun:"type:uuid" json:"userId"`
	Moved_in  bool      `json:"movedIn"`
	Moved_out bool      `json:"movedOut"`
}

type ApplicationsCars []*ApplicationCar

func (item *ApplicationCar) SetForeignKeys() {
	item.UserID = item.User.ID
}
