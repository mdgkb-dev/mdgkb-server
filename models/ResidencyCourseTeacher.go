package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ResidencyCourseTeacher struct {
	bun.BaseModel     `bun:"residency_courses_teachers,alias:postgraduate_courses_teachers"`
	ID                uuid.UUID        `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	ResidencyCourse   *ResidencyCourse `bun:"rel:belongs-to" json:"residencyCourse"`
	ResidencyCourseID uuid.NullUUID    `bun:"type:uuid" json:"residencyCourseId"`
	Teacher           *Teacher         `bun:"rel:belongs-to" json:"teacher"`
	TeacherID         uuid.UUID        `bun:"type:uuid" json:"teacherId"`
	Main              bool             `json:"main"`
}

type ResidencyCoursesTeachers []*ResidencyCourseTeacher
