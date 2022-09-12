package models

import (
	"github.com/google/uuid"
)

type ScheduleItem struct {
	ID             uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	ScheduleID     uuid.UUID     `bun:"type:uuid" json:"scheduleId"`
	TimetableDayID uuid.UUID     `bun:"type:uuid" json:"timetableDayId"`
	Name           string        `json:"name"`
	StartTime      string        `json:"startTime"`
	EndTime        string        `json:"endTime"`
	Dishes         Dishes        `bun:"rel:has-many" json:"dishes"`
}

type ScheduleItems []*ScheduleItem
