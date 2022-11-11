package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Regalia struct {
	bun.BaseModel `bun:"regalias,alias:regalias"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`

	EmployeeID uuid.NullUUID `bun:"type:uuid" json:"employeeId"`
	Employee   *Employee     `bun:"rel:belongs-to" json:"employee"`
}

type Regalias []*Regalia
