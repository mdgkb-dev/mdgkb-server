package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PostgraduateCourseTeacher struct {
	bun.BaseModel        `bun:"postgraduate_courses_teachers,alias:postgraduate_courses_teachers"`
	ID                   uuid.UUID           `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	PostgraduateCourse   *PostgraduateCourse `bun:"rel:belongs-to" json:"postgraduateCourse"`
	PostgraduateCourseID uuid.NullUUID       `bun:"type:uuid" json:"postgraduateCourseId"`
	Teacher              *Teacher            `bun:"rel:belongs-to" json:"teacher"`
	TeacherID            uuid.UUID           `bun:"type:uuid" json:"teacherId"`
	Main                 bool                `json:"main"`
}

type PostgraduateCoursesTeachers []*PostgraduateCourseTeacher
