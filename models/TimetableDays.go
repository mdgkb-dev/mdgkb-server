package models

import (
	"github.com/google/uuid"
	"time"
)

type TimetableDay struct {
	ID          uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	TimetableID uuid.UUID `bun:"type:uuid" json:"timetableId"`
	Weekday     *Weekday  `bun:"rel:belongs-to" json:"weekday"`
	WeekdayId   uuid.UUID `bun:"type:uuid" json:"weekdayId"`
	StartTime   time.Time `bun:"type:time" json:"startTime"`
	EndTime     time.Time `bun:"type:time" json:"endTime"`
	BreakExist  bool      `json:"breakExist"`
	BreakStart  time.Time `bun:"type:time" json:"breakStartTime"`
	BreakEnd    time.Time `bun:"type:time" json:"breakEndTime"`
}
