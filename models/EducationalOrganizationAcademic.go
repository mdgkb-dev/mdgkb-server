package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type EducationalOrganizationAcademic struct {
	bun.BaseModel `bun:"educational_organization_academics,select:educational_organization_academics_view,alias:educational_organization_academics_view"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	DoctorID      uuid.NullUUID `bun:"type:uuid" json:"doctorId"`
	Doctor        *Doctor       `bun:"rel:belongs-to" json:"doctor"`
}

type EducationalOrganizationAcademics []*EducationalOrganizationAcademic
