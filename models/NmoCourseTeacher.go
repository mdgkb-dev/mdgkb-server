package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type NmoCourseTeacher struct {
	bun.BaseModel `bun:"nmo_courses_teachers,alias:nmo_courses_teachers"`
	ID            uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	NmoCourse     *NmoCourse    `bun:"rel:belongs-to" json:"nmoCourse"`
	NmoCourseID   uuid.NullUUID `bun:"type:uuid" json:"nmoCourseId"`
	Teacher       *Teacher      `bun:"rel:belongs-to" json:"teacher"`
	TeacherID     uuid.UUID     `bun:"type:uuid" json:"teacherId"`
	Main          bool          `json:"main"`
}

type NmoCoursesTeachers []*NmoCourseTeacher
