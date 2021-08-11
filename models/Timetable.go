package models

import (
	"github.com/google/uuid"
)

type Timetable struct {
	ID                     uuid.UUID       `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Description            string          `json:"description"`
	TimetableDays          []*TimetableDay `bun:"rel:has-many" json:"timetableDays"`
	TimetableDaysForDelete []string        `bun:"-" json:"timetableDaysForDelete"`
}

func (timetable *Timetable) SetIdForChildren() {
	if len(timetable.TimetableDays) < 0 {
		return
	}
	for i := range timetable.TimetableDays {
		timetable.TimetableDays[i].TimetableID = timetable.ID
	}
}
