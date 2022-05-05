package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type DpoBaseCourse struct {
	bun.BaseModel `bun:"dpo_base_courses,alias:dpo_base_courses"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Order         int       `bun:"dpo_base_course_order" json:"order"`
	Cost          int       `json:"cost"`
	Start         time.Time `bun:"dpo_base_course_start" json:"start"`
	Listeners     int       `json:"listeners"`
	Hours         int       `json:"hours"`
	TeacherID     uuid.UUID `bun:"type:uuid" json:"teacherId"`
	Teacher       *Teacher  `bun:"rel:belongs-to" json:"teacher"`
}

type DpoBaseCourses []*DpoBaseCourse
