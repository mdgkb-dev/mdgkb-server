package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type NmoCourseDates struct {
	bun.BaseModel `bun:"nmo_courses_dates,alias:nmo_courses_dates"`
	ID            uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Start         time.Time     `bun:"dpo_course_start" json:"start"`
	End           time.Time     `bun:"dpo_course_end" json:"end"`
	NmoCourse     *NmoCourse    `bun:"rel:belongs-to" json:"nmoCourse"`
	NmoCourseID   uuid.NullUUID `bun:"type:uuid" json:"nmoCourseId"`
}

type NmoCoursesDates []*NmoCourseDates
