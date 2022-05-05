package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type EducationalManager struct {
	bun.BaseModel `bun:"educational_managers,select:educational_managers_view,alias:educational_managers_view"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	DoctorID      uuid.UUID `bun:"type:uuid" json:"doctorId"`
	Doctor        *Doctor   `bun:"rel:belongs-to" json:"doctor"`
	Role          string    `json:"role"`
	Order         int       `bun:"educational_manager_order" json:"order"`
}

type EducationalManagers []*EducationalManager
