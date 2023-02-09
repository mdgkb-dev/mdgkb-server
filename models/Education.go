package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Education struct {
	bun.BaseModel  `bun:"educations,alias:educations"`
	ID             uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Type           string    `json:"type"`
	Institution    string    `json:"institution"`
	Document       string    `json:"document"`
	Qualification  string    `json:"qualification"`
	Specialization string    `json:"specialization"`
	Start          time.Time `bun:"education_start" json:"start"`
	End            time.Time `bun:"education_end" json:"end"`

	EmployeeID uuid.NullUUID `bun:"type:uuid" json:"employeeId"`
	Employee   *Employee     `bun:"rel:belongs-to" json:"employee"`
}

type Educations []*Education
