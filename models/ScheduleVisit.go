package models

import (
	"github.com/google/uuid"
)

type ScheduleVisit struct {
	ID                          uuid.UUID            `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                        string               `json:"name"`
	Description                 string               `json:"description"`
	ScheduleVisitItems          []*ScheduleVisitItem `bun:"rel:has-many" json:"scheduleVisitItems"`
	ScheduleVisitItemsForDelete []string             `bun:"-" json:"scheduleVisitItemsForDelete"`
}

func (item *ScheduleVisit) SetIdForChildren() {
	if len(item.ScheduleVisitItems) == 0 {
		return
	}
	for i := range item.ScheduleVisitItems {
		item.ScheduleVisitItems[i].ScheduleVisitID = item.ID
	}
}
