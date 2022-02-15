package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MedicalProfile struct {
	bun.BaseModel            `bun:"medical_profiles,alias:medical_profiles"`
	ID                       uuid.UUID                `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                     string                   `json:"name"`
	Description              string                   `json:"description"`
	Icon                     *FileInfo                `bun:"rel:belongs-to" json:"icon"`
	IconId                   uuid.NullUUID            `bun:"type:uuid" json:"iconId"`
	MedicalProfilesDivisions MedicalProfilesDivisions `bun:"rel:has-many" json:"medicalProfilesDivisions"`
	MedicalProfilesNews      MedicalProfilesNews      `bun:"rel:has-many" json:"medicalProfilesNews"`
}

type MedicalProfiles []*MedicalProfile
