package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type NmoCourseSpecialization struct {
	bun.BaseModel    `bun:"nmo_courses_specializations,alias:nmo_courses_specializations"`
	ID               uuid.UUID       `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	NmoCourse        *NmoCourse      `bun:"rel:belongs-to" json:"nmoCourse"`
	NmoCourseID      uuid.NullUUID   `bun:"type:uuid" json:"nmoCourseId"`
	Specialization   *Specialization `bun:"rel:belongs-to" json:"specialization"`
	SpecializationID uuid.UUID       `bun:"type:uuid" json:"specializationId"`
}

type NmoCoursesSpecializations []*NmoCourseSpecialization
