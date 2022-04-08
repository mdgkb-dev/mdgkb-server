package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ResidencyCourseSpecialization struct {
	bun.BaseModel     `bun:"residency_courses_specializations,alias:residency_courses_specializations"`
	ID                uuid.UUID        `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Main              bool             `json:"main"`
	ResidencyCourse   *ResidencyCourse `bun:"rel:belongs-to" json:"residencyCourse"`
	ResidencyCourseID uuid.NullUUID    `bun:"type:uuid" json:"residencyCourseId"`
	Specialization    *Specialization  `bun:"rel:belongs-to" json:"specialization"`
	SpecializationID  uuid.UUID        `bun:"type:uuid" json:"specializationId"`
}

type ResidencyCoursesSpecializations []*ResidencyCourseSpecialization
