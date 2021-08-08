package models

import "github.com/google/uuid"

type Timetable struct {
	ID            uuid.UUID       `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	TimetableDays []*TimetableDay `bun:"rel:has-many" json:"timetableDays"`
}
