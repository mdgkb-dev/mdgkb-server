package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type EducationalOrganizationTeacher struct {
	bun.BaseModel `bun:"educational_organization_teachers,alias:educational_organization_teachers"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	DoctorID      uuid.UUID `bun:"type:uuid" json:"doctorId"`
	Doctor        *Doctor   `bun:"rel:belongs-to" json:"doctor"`
	Position      string    `json:"position"`
}

type EducationalOrganizationTeachers []*EducationalOrganizationTeacher
