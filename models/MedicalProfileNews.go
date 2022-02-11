package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type MedicalProfileNews struct {
	bun.BaseModel    `bun:"medical_profiles_news,alias:medical_profiles_news"`
	ID               uuid.UUID       `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	NewsID           uuid.NullUUID   `bun:"type:uuid" json:"newsId"`
	News             *News           `bun:"rel:belongs-to" json:"news"`
	MedicalProfileID uuid.UUID       `bun:"type:uuid" json:"medicalProfileId"`
	MedicalProfile   *MedicalProfile `bun:"rel:belongs-to" json:"medicalProfile"`
}

type MedicalProfilesNews []*MedicalProfileNews
