package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MedicalProfileDivision struct {
	bun.BaseModel    `bun:"medical_profiles_divisions,alias:medical_profiles_divisions"`
	ID               uuid.NullUUID   `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	DivisionID       uuid.NullUUID   `bun:"type:uuid" json:"divisionId"`
	Division         *Division       `bun:"rel:belongs-to" json:"division"`
	MedicalProfileID uuid.NullUUID   `bun:"type:uuid" json:"medicalProfileId"`
	MedicalProfile   *MedicalProfile `bun:"rel:belongs-to" json:"medicalProfile"`
}

type MedicalProfilesDivisions []*MedicalProfileDivision
