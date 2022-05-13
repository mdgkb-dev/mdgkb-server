package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type CandidateApplicationSpecialization struct {
	bun.BaseModel          `bun:"candidate_application_specializations,alias:candidate_application_specializations"`
	ID                     uuid.UUID             `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	CandidateApplication   *CandidateApplication `bun:"rel:belongs-to" json:"candidateApplication"`
	CandidateApplicationID uuid.NullUUID         `bun:"type:uuid" json:"candidateApplicationId"`
	Specialization         *Specialization       `bun:"rel:belongs-to" json:"specialization"`
	SpecializationID       uuid.UUID             `bun:"type:uuid" json:"specializationId"`
}

type CandidateApplicationSpecializations []*CandidateApplicationSpecialization
