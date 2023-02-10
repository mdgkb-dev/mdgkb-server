package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Accreditation struct {
	bun.BaseModel `bun:"accreditations,alias:accreditations"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Document      string        `json:"document"`

	Specialization string    `json:"specialization"`
	StartDate      time.Time `json:"startDate"`
	EndDate        time.Time `json:"endDate"`

	EmployeeID uuid.NullUUID `bun:"type:uuid" json:"employeeId"`
	Employee   *Employee     `bun:"rel:belongs-to" json:"employee"`
}

type Accreditations []*Accreditation
