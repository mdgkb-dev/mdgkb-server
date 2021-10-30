package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Event struct {
	bun.BaseModel     `bun:"events,alias:events"`
	ID                uuid.NullUUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	EventApplications EventApplications `bun:"rel:has-many" json:"eventApplications"`
}

type Events []*Event
