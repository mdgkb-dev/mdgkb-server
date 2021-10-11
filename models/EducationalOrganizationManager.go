package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type EducationalOrganizationManager struct {
	bun.BaseModel `bun:"educational_organization_managers,alias:educational_organization_managers"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	DoctorID      uuid.UUID `bun:"type:uuid" json:"doctorId"`
	Doctor        *Doctor   `bun:"rel:belongs-to" json:"doctor"`
	Role          *string   `json:"role"`
	ManagerOrder  *int      `json:"managerOrder"`
}

type EducationalOrganizationManagers []*EducationalOrganizationManager
