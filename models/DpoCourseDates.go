package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DpoCourseDates struct {
	bun.BaseModel `bun:"dpo_courses_dates,alias:dpo_courses_dates"`
	ID            uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Start         time.Time     `bun:"dpo_course_start" json:"start"`
	End           time.Time     `bun:"dpo_course_end" json:"end"`
	DpoCourse     *DpoCourse    `bun:"rel:belongs-to" json:"dpoCourse"`
	DpoCourseID   uuid.NullUUID `bun:"type:uuid" json:"dpoCourseId"`
}

type DpoCoursesDates []*DpoCourseDates
