package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DpoCourseSpecialization struct {
	bun.BaseModel    `bun:"dpo_courses_specializations,alias:dpo_courses_specializations"`
	ID               uuid.UUID       `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	DpoCourse        *DpoCourse      `bun:"rel:belongs-to" json:"dpoCourse"`
	DpoCourseID      uuid.NullUUID   `bun:"type:uuid" json:"dpoCourseId"`
	Specialization   *Specialization `bun:"rel:belongs-to" json:"specialization"`
	SpecializationID uuid.UUID       `bun:"type:uuid" json:"specializationId"`
}

type DpoCoursesSpecializations []*DpoCourseSpecialization
