package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type EventApplication struct {
	bun.BaseModel `bun:"event_applications,alias:event_applications"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `

	Event   *Event        `bun:"rel:belongs-to" json:"event"`
	EventID uuid.NullUUID `bun:"type:uuid" json:"eventId"`

	User   *User         `bun:"rel:belongs-to" json:"user"`
	UserID uuid.NullUUID `bun:"type:uuid" json:"userId"`
}

type EventApplications []*EventApplication
