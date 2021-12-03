package models

import (
	"github.com/google/uuid"
)

type TimetableDay struct {
	ID             uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	IsWeekend      bool      `json:"isWeekend"`
	StartTime      string    `json:"startTime"`
	EndTime        string    `json:"endTime"`
	BreaksExists   bool      `json:"breaksExists"`
	AroundTheClock bool      `json:"aroundTheClock"`

	Weekday               *Weekday          `bun:"rel:belongs-to" json:"weekday"`
	WeekdayId             uuid.UUID         `bun:",nullzero,notnull,type:uuid" json:"weekdayId"`
	TimetablePattern      *TimetablePattern `bun:"rel:belongs-to" json:"timetablePattern"`
	TimetablePatternID    uuid.UUID         `bun:"type:uuid,nullzero,default:NULL" json:"timetablePatternId"`
	Timetable             *Timetable        `bun:"rel:belongs-to" json:"timetable"`
	TimetableID           uuid.UUID         `bun:"type:uuid,nullzero,default:NULL" json:"timetableId"`
	BreakPeriods          TimePeriods       `bun:"rel:has-many" json:"breakPeriods"`
	BreakPeriodsForDelete []string          `bun:"-" json:"breakPeriodsForDelete"`
}

type TimetableDays []*TimetableDay

func (item *TimetableDay) SetIdForChildren() {
	if len(item.BreakPeriods) == 0 {
		return
	}
	for i := range item.BreakPeriods {
		item.BreakPeriods[i].TimetableDayID = item.ID
	}
}

func (items TimetableDays) SetIdForChildren() {
	for i := range items {
		items[i].SetIdForChildren()
	}
}

func (items TimetableDays) GetIDForDelete() []string {
	idPool := make([]string, 0)
	for _, item := range items {
		idPool = append(idPool, item.BreakPeriodsForDelete...)
	}
	return idPool
}

func (items TimetableDays) GetTimePeriods() TimePeriods {
	itemsForGet := make(TimePeriods, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.BreakPeriods...)
	}
	return itemsForGet
}

func (items TimetableDays) SetForeignKeys() {
	for i := range items {
		if items[i].Timetable != nil {
			items[i].TimetableID = items[i].Timetable.ID
		}
		if items[i].TimetablePattern != nil {
			items[i].TimetablePatternID = items[i].TimetablePattern.ID
		}
	}
}
