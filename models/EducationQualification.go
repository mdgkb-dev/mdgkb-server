package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type EducationQualification struct {
	bun.BaseModel `bun:"education_qualifications,alias:education_qualifications"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Document      string    `json:"document"`

	Specialization string    `json:"specialization"`
	StartDate      time.Time `json:"startDate"`
	EndDate        time.Time `json:"endDate"`
}

type EducationQualifications []*EducationQualification
