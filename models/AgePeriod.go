package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type AgePeriod struct {
	bun.BaseModel `bun:"age_periods,alias:age_periods"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
}

type AgePeriods []*AgePeriod
