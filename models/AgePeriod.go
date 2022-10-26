package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DietAge struct {
	bun.BaseModel `bun:"diet_ages,alias:diet_ages"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Timetable     *Timetable    `bun:"rel:belongs-to" json:"timetable"`
	TimetableID   uuid.NullUUID `bun:"type:uuid"  json:"timetableId"`
	Diet          *Diet         `bun:"rel:has-one,join:diet_id=id" json:"diet"`
	DietID        uuid.NullUUID `bun:"type:uuid"  json:"dietId"`
	Order         uint          `bun:"diet_age_order" json:"order"`
}

type DietAges []*DietAge

func (item *DietAge) SetForeignKeys() {
	if item.Timetable != nil {
		item.TimetableID = item.Timetable.ID
	}
}
