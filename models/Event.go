package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Event struct {
	bun.BaseModel     `bun:"events,alias:events"`
	ID                uuid.NullUUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	EventApplications EventApplications `bun:"rel:has-many" json:"eventApplications"`

	News *News `bun:"rel:has-one" json:"news"`

	Form   *Form     `bun:"rel:belongs-to" json:"form"`
	FormID uuid.UUID `bun:"type:uuid" json:"formId"`
}

type Events []*Event

func (item *Event) SetForeignKeys() {
	item.FormID = item.Form.ID
}
