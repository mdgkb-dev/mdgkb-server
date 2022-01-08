package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type EducationalOrganizationAcademic struct {
	bun.BaseModel `bun:"educational_organization_academics,alias:educational_organization_academics"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	DoctorID      uuid.UUID `bun:"type:uuid" json:"doctorId"`
	Doctor        *Doctor   `bun:"rel:belongs-to" json:"doctor"`
}

type EducationalOrganizationAcademics []*EducationalOrganizationAcademic
