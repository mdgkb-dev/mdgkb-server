package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type TeachingActivity struct {
	bun.BaseModel `bun:"teaching_activities,alias:teaching_activities"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`

	DoctorID uuid.NullUUID `bun:"type:uuid" json:"doctorId"`
	Doctor   *Doctor       `bun:"rel:belongs-to" json:"doctor"`
}

type TeachingActivities []*TeachingActivity
