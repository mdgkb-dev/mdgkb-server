package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MedicalProfile struct {
	bun.BaseModel `bun:"medical_profiles,alias:medical_profiles"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`
}

type MedicalProfiles []*MedicalProfile
