package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PostgraduateCourseSpecialization struct {
	bun.BaseModel        `bun:"postgraduate_courses_specializations,alias:postgraduate_courses_specializations"`
	ID                   uuid.UUID           `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	PostgraduateCourse   *PostgraduateCourse `bun:"rel:belongs-to" json:"postgraduateCourse"`
	PostgraduateCourseID uuid.NullUUID       `bun:"type:uuid" json:"postgraduateCourseId"`
	Specialization       *Specialization     `bun:"rel:belongs-to" json:"specialization"`
	SpecializationID     uuid.UUID           `bun:"type:uuid" json:"specializationId"`
	Main                 bool                `json:"main"`
}

type PostgraduateCoursesSpecializations []*PostgraduateCourseSpecialization
