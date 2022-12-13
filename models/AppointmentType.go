package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type AppointmentType struct {
	bun.BaseModel `bun:"appointments_types,alias:appointments_types"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name          string        `json:"name"`
	Order         uint8         `bun:"item_order" json:"order"`
	Description   string        `json:"description"`
	OMS           bool          `json:"oms"`
	FormPattern   *FormPattern  `bun:"rel:belongs-to" json:"formPattern"`
	FormPatternID uuid.NullUUID `bun:"type:uuid" json:"formPatternId"`
}

type AppointmentsTypes []*AppointmentType
