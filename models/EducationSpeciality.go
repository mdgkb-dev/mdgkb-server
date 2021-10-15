package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type EducationSpeciality struct {
	bun.BaseModel `bun:"education_specialities,alias:education_specialities"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`
}

type EducationSpecialities []*EducationSpeciality
