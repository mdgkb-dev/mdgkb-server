package models

import (
	"github.com/google/uuid"
)

type Schedule struct {
	ID                     uuid.UUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                   string        `json:"name"`
	Description            string        `json:"description"`
	ScheduleItems          ScheduleItems `bun:"rel:has-many" json:"scheduleItems"`
	ScheduleItemsForDelete []string      `bun:"-" json:"scheduleItemsForDelete"`
}

func (item *Schedule) SetIdForChildren() {
	if len(item.ScheduleItems) == 0 {
		return
	}
	for i := range item.ScheduleItems {
		item.ScheduleItems[i].ScheduleID = item.ID
	}
}
