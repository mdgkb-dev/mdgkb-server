package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Diet struct {
	bun.BaseModel `bun:"diets,alias:diets"`
	ID            uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	ShortName     string        `json:"shortName"`
	SiteName      string        `json:"siteName"`
	Diabetes      bool          `json:"diabetes"`
	MotherDiet    *Diet         `bun:"rel:belongs-to" json:"motherDiet"`
	MotherDietID  uuid.NullUUID `bun:"type:uuid"  json:"motherDietId"`
	AgePeriod     *AgePeriod    `bun:"rel:belongs-to" json:"agePeriod"`
	AgePeriodID   uuid.NullUUID `bun:"type:uuid"  json:"agePeriodId"`
	Timetable     *Timetable    `bun:"rel:belongs-to" json:"timetable"`
	TimetableID   uuid.NullUUID `bun:"type:uuid"  json:"timetableId"`
}

type Diets []*Diet

func (item *Diet) SetForeignKeys() {
	if item.Timetable != nil {
		item.TimetableID = item.Timetable.ID
	}
	if item.AgePeriod != nil {
		item.AgePeriodID = item.AgePeriod.ID
	}
}