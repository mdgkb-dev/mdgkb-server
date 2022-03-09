package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DpoCourseTeacher struct {
	bun.BaseModel `bun:"dpo_courses_teachers,alias:dpo_courses_teachers"`
	ID            uuid.UUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	DpoCourse     *DpoCourse    `bun:"rel:belongs-to" json:"dpoCourse"`
	DpoCourseID   uuid.NullUUID `bun:"type:uuid" json:"dpoCourseId"`
	Teacher       *Teacher      `bun:"rel:belongs-to" json:"teacher"`
	TeacherID     uuid.UUID     `bun:"type:uuid" json:"teacherId"`
	Main          bool          `json:"main"`
}

type DpoCoursesTeachers []*DpoCourseTeacher
