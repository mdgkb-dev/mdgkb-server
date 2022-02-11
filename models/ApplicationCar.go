package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type ApplicationCar struct {
	bun.BaseModel `bun:"banners,alias:banners"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	CarNumber     string        `json:"carNumber"`
	CarBrand      string        `json:"carBrand"`

	Division   *Division     `bun:"rel:belongs-to" json:"division"`
	DivisionId uuid.NullUUID `bun:"type:uuid" json:"divisionId,omitempty"`

	Date time.Time `bun:"hospitalization_date" json:"date"`

	User   *User     `bun:"rel:belongs-to" json:"user"`
	UserID uuid.UUID `bun:"type:uuid" json:"userId"`
}

type ApplicationsCars []*ApplicationCar

func (item *ApplicationCar) SetForeignKeys() {
	item.UserID = item.User.ID
}
