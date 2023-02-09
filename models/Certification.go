package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Certification struct {
	bun.BaseModel     `bun:"certifications,alias:certifications"`
	ID                uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Specialization    string        `json:"specialization"`
	CertificationDate time.Time     `json:"certificationDate"`
	EndDate           time.Time     `json:"endDate"`
	Place             string        `json:"place"`
	Document          string        `json:"document"`

	EmployeeID uuid.NullUUID `bun:"type:uuid" json:"employeeId"`
	Employee   *Employee     `bun:"rel:belongs-to" json:"employee"`
}

type Certifications []*Certification
