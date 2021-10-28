package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Event struct {
	bun.BaseModel `bun:"events,alias:events"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `

	News   News          `bun:"rel:belongs-to" json:"news"`
	NewsID uuid.NullUUID `bun:"type:uuid" json:"newsId"`

	EventApplications EventApplications `bun:"rel:has-many" json:"eventApplications"`
}

type Events []*Event
