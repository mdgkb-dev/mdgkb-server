package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type TeachingActivity struct {
	bun.BaseModel `bun:"teaching_activities,alias:teaching_activities"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`

	EmployeeID uuid.NullUUID `bun:"type:uuid" json:"employeeId"`
	Employee   *Employee     `bun:"rel:belongs-to" json:"employee"`
}

type TeachingActivities []*TeachingActivity
