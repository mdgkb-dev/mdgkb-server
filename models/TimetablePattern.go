package models

import "github.com/uptrace/bun"

type TimetablePattern struct {
	bun.BaseModel `bun:"timetable_patterns,alias:timetable_patterns"`
	Title         string `json:"title"`
	Timetable
}

type TimetablePatterns []*TimetablePattern

func (item *TimetablePattern) SetIDForChildren() {
	if len(item.TimetableDays) == 0 {
		return
	}
	for i := range item.TimetableDays {
		item.TimetableDays[i].TimetablePatternID = item.ID
	}
}
