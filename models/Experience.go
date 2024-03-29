package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Experience struct {
	bun.BaseModel `bun:"experiences,alias:experiences"`
	ID            uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	EmployeeID    uuid.NullUUID `bun:"type:uuid" json:"employeeId"`
	Employee      *Employee     `bun:"rel:belongs-to" json:"employee"`
	Start         time.Time     `bun:"item_start" json:"start"`
	End           time.Time     `bun:"item_end" json:"end"`
	Place         string        `json:"place"`
	Position      string        `json:"position"`
}

type Experiences []*Experience
