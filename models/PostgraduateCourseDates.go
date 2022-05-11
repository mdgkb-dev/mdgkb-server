package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PostgraduateCourseDates struct {
	bun.BaseModel        `bun:"postgraduate_courses_dates,alias:postgraduate_courses_dates"`
	ID                   uuid.UUID           `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Start                time.Time           `bun:"postgraduate_course_start" json:"start"`
	End                  time.Time           `bun:"postgraduate_course_end" json:"end"`
	PostgraduateCourse   *PostgraduateCourse `bun:"rel:belongs-to" json:"postgraduateCourse"`
	PostgraduateCourseID uuid.NullUUID       `bun:"type:uuid" json:"postgraduateCourseId"`
}

type PostgraduateCoursesDates []*PostgraduateCourseDates
