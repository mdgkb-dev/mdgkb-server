package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DpoAppointment struct {
	bun.BaseModel `bun:"dpo_appointments,alias:dpo_appointments"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `

	Application   *FileInfo      `json:"application"`
	ApplicationID *uuid.NullUUID `json:"applicationId"`

	DpoCourseID uuid.NullUUID `bun:"type:uuid" json:"dpoCourseId"`
	DpoCourse   *DpoCourse    `bun:"rel:belongs-to" json:"dpoCourse"`

	User   *User     `bun:"rel:belongs-to" json:"user"`
	UserID uuid.UUID `bun:"type:uuid" json:"userId"`
}

type DpoAppointments []*DpoAppointment

func (item *DpoAppointment) SetForeignKeys() {
	item.UserID = item.User.ID
}
