package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Appointment struct {
	bun.BaseModel `bun:"appointments,select:appointments_view,alias:appointments_view"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	Date time.Time `bun:"item_date" json:"date"`
	Time string    `bun:"item_time" json:"time"`

	Specialization   *Specialization `bun:"rel:belongs-to" json:"specialization"`
	SpecializationID uuid.NullUUID   `bun:"type:uuid" json:"specializationId"`

	DoctorID uuid.NullUUID `bun:"type:uuid" json:"doctorId"`
	Doctor   *Doctor       `bun:"rel:belongs-to" json:"doctor"`

	AppointmentTypeID uuid.NullUUID    `bun:"type:uuid" json:"appointmentTypeId"`
	AppointmentType   *AppointmentType `bun:"rel:belongs-to" json:"appointmentType"`

	FormValue   *FormValue    `bun:"rel:belongs-to" json:"formValue"`
	FormValueID uuid.NullUUID `bun:"type:uuid" json:"formValueId"`
}

type Appointments []*Appointment

func (item *Appointment) SetForeignKeys() {
	if item.FormValue != nil {
		item.FormValueID = item.FormValue.ID
	}
}

func (item *Appointment) SetFilePath(fileID *string) *string {
	path := item.FormValue.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}
