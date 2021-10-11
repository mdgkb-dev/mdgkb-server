package models

import (
	"github.com/google/uuid"
)

type TimetableDay struct {
	ID             uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	IsWeekend      bool      `json:"isWeekend"`
	TimetableID    uuid.UUID `bun:"type:uuid" json:"timetableId"`
	Weekday        *Weekday  `bun:"rel:belongs-to" json:"weekday"`
	WeekdayId      uuid.UUID `bun:",nullzero,notnull,type:uuid" json:"weekdayId"`
	StartTime      *string   ` json:"startTime"`
	EndTime        *string   ` json:"endTime"`
	BreakExist     bool      `json:"breakExist"`
	BreakStartTime *string   ` json:"breakStartTime"`
	BreakEndTime   *string   `json:"breakEndTime"`
	IsCustom       bool      `json:"isCustom"`
	CustomName     string    `json:"customName"`
}

type TimetableDays []*TimetableDay
