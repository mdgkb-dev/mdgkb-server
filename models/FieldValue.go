package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type FieldValue struct {
	bun.BaseModel `bun:"field_values,alias:field_values"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	ValueString   string    `json:"valueString"`
	ValueNumber   int       `json:"valueNumber"`
	ValueDate     time.Time `json:"valueDate"`

	Field   *Field `bun:"rel:belongs-to" json:"field"`
	FieldID uuid.UUID          `bun:"type:uuid" json:"fieldId"`

	EventApplication   *EventApplication `bun:"rel:belongs-to" json:"eventApplication"`
	EventApplicationID uuid.UUID          `bun:"type:uuid" json:"eventApplicationId"`
}

type FieldValues []*FieldValue
