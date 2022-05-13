package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type TimePeriod struct {
	bun.BaseModel  `bun:"time_periods,alias:time_periods"`
	ID             uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	StartTime      string    `json:"startTime"`
	EndTime        string    `json:"endTime"`
	TimetableDayID uuid.UUID `bun:"type:uuid" json:"timetableDayId"`
}

type TimePeriods []*TimePeriod
