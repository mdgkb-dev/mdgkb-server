package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type DpoCourse struct {
	bun.BaseModel `bun:"dpo_courses,alias:dpo_courses"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Order         int       `bun:"dpo_course_order" json:"order"`
	Start         time.Time `bun:"dpo_course_start" json:"start"`
	Listeners     int       `json:"listeners"`
	Hours         int       `json:"hours"`
	TeacherID     uuid.UUID `bun:"type:uuid" json:"teacherId"`
	Teacher       *Teacher  `bun:"rel:belongs-to" json:"teacher"`
}

type DpoCourses []*DpoCourse
