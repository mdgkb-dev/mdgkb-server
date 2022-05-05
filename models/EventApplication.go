package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type EventApplication struct {
	bun.BaseModel `bun:"event_applications,alias:event_applications"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	Event   *Event        `bun:"rel:belongs-to" json:"event"`
	EventID uuid.NullUUID `bun:"type:uuid" json:"eventId"`

	User   *User     `bun:"rel:belongs-to" json:"user"`
	UserID uuid.UUID `bun:"type:uuid" json:"userId"`

	FieldValues FieldValues `bun:"rel:has-many" json:"fieldValues"`
}

type EventApplications []*EventApplication

func (item *EventApplication) SetIdForChildren() {
	for i := range item.FieldValues {
		item.FieldValues[i].EventApplicationID = item.ID
	}
}
