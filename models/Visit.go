package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Visit struct {
	bun.BaseModel `bun:"visits,select:visits,alias:visits"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Date          *time.Time    `bun:",nullzero" json:"date"`
	Entered       bool          `json:"entered"`
	Exited        bool          `json:"exited"`

	ApplicationCar   *ApplicationCar `bun:"rel:belongs-to" json:"applicationCar"`
	ApplicationCarID uuid.NullUUID   `bun:"type:uuid" json:"applicationCarId,omitempty"`
}

type Visits []*Visit
