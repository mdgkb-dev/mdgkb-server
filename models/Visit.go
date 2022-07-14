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

	VisitsApplication   *VisitsApplication `bun:"rel:belongs-to" json:"visitsApplication"`
	VisitsApplicationID uuid.NullUUID      `bun:"type:uuid" json:"visitsApplicationId,omitempty"`
}

type Visits []*Visit
