package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type CandidateExam struct {
	bun.BaseModel                                `bun:"candidate_exams,alias:candidate_exams"`
	ID                                           uuid.NullUUID                       `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	CandidateApplicationSpecializations          CandidateApplicationSpecializations `bun:"rel:has-many" json:"candidateApplicationSpecializations"`
	CandidateApplicationSpecializationsForDelete []uuid.UUID                         `bun:"-" json:"candidateApplicationSpecializationsForDelete"`

	EducationYear   *EducationYear `bun:"rel:belongs-to" json:"educationYear"`
	EducationYearID uuid.UUID      `bun:"type:uuid" json:"educationYearId"`

	FormPattern   *FormPattern  `bun:"rel:belongs-to" json:"formPattern"`
	FormPatternID uuid.NullUUID `bun:"type:uuid" json:"formPatternId"`

	DocumentType   *DocumentType `bun:"rel:belongs-to" json:"documentType"`
	DocumentTypeID uuid.NullUUID `bun:"type:uuid" json:"documentTypeId"`
}

type CandidateExams []*CandidateExam
