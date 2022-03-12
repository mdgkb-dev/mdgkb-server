package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Teacher struct {
	bun.BaseModel `bun:"teachers_view,alias:teachers_view"`
	ID            uuid.UUID  `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	DoctorID      uuid.UUID  `bun:"type:uuid" json:"doctorId"`
	Doctor        *Doctor    `bun:"rel:belongs-to" json:"doctor"`
	Position      string     `json:"position"`
	DpoCourses    DpoCourses `bun:"rel:has-many" json:"dpoCourses"`
}

type Teachers []*Teacher