package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MedicalProfile struct {
	bun.BaseModel            `bun:"medical_profiles,alias:medical_profiles"`
	ID                       uuid.NullUUID            `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                     string                   `json:"name"`
	Description              string                   `json:"description"`
	Icon                     string                   `json:"icon"`
	IconID                   uuid.NullUUID            `bun:"type:uuid" json:"iconId"`
	SvgCode                  string                   `json:"svgCode"`
	MedicalProfilesDivisions MedicalProfilesDivisions `bun:"rel:has-many" json:"medicalProfilesDivisions"`
	MedicalProfilesNews      MedicalProfilesNews      `bun:"rel:has-many" json:"medicalProfilesNews"`
}

type MedicalProfiles []*MedicalProfile
