package models

import (
	"fmt"
	"github.com/google/uuid"
)

type Timetable struct {
	ID                     uuid.UUID       `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Description            string          `json:"description"`
	TimetableDays          []*TimetableDay `bun:"rel:has-many" json:"timetableDays"`
	TimetableDaysForDelete []string        `bun:"-" json:"timetableDaysForDelete"`
}

func (timetable *Timetable) SetIdForChildren() {
	for i := range timetable.TimetableDays {
		timetable.TimetableDays[i].TimetableID = timetable.ID
	}
	fmt.Println(timetable.TimetableDays[0])
}
