package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type EducationAccreditation struct {
	bun.BaseModel `bun:"education_accreditations,alias:education_accreditations"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Document      string        `json:"document"`

	Specialization string    `json:"specialization"`
	StartDate      time.Time `json:"startDate"`
	EndDate        time.Time `json:"endDate"`
}

type EducationAccreditations []*EducationAccreditation
