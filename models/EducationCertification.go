package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type EducationCertification struct {
	bun.BaseModel     `bun:"education_certifications,alias:education_certifications"`
	ID                uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Specialization    string        `json:"specialization"`
	CertificationDate time.Time     `json:"certificationDate"`
	EndDate           time.Time     `json:"endDate"`
	Place             string        `json:"place"`
	Document          string        `json:"document"`
}

type EducationCertifications []*EducationCertification
