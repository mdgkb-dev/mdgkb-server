package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Dish struct {
	bun.BaseModel  `bun:"dishes,alias:dishes"`
	ID             uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name           string        `json:"name"`
	Weight         string        `json:"weight"`
	ScheduleItem   *ScheduleItem `bun:"rel:belongs-to" json:"scheduleItem"`
	ScheduleItemID uuid.NullUUID `bun:"type:uuid" json:"scheduleItemId"`
}

type Dishes []*Dish
