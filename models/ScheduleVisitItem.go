package models

import (
	"github.com/google/uuid"
)

type ScheduleVisitItem struct {
	ID              uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	ScheduleVisitID uuid.UUID `bun:"type:uuid" json:"ScheduleVisitId"`
	Name            string    `json:"name"`
	StartTime       *string   `json:"startTime"`
	EndTime         *string   `json:"endTime"`
}
