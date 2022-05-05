package models

import (
	"github.com/google/uuid"
)

type Timetable struct {
	ID                     uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Description            string        `json:"description"`
	TimetableDays          TimetableDays `bun:"rel:has-many" json:"timetableDays"`
	TimetableDaysForDelete []string      `bun:"-" json:"timetableDaysForDelete"`
}

func (item *Timetable) SetIdForChildren() {
	if len(item.TimetableDays) == 0 {
		return
	}
	for i := range item.TimetableDays {
		item.TimetableDays[i].TimetableID = item.ID
	}
}

func (item *Timetable) InitAppointmentsSlots() {
	item.TimetableDays.InitAppointmentsSlots()
}
