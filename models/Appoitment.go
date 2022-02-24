package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Appointment struct {
	bun.BaseModel `bun:"appointments,alias:appointments"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `

	Date time.Time `bun:"appointment_date" json:"date"`
	Time string    `bun:"appointment_time" json:"time"`

	ClinicName           string `json:"clinicName"`
	ClinicReferralNumber string `json:"clinicReferralNumber"`

	FormScan   *FileInfo     `bun:"rel:belongs-to" json:"formScan"`
	FormScanID uuid.NullUUID `bun:"type:uuid" json:"formScanId"`

	Specialization   *Specialization `bun:"rel:belongs-to" json:"specialization"`
	SpecializationID uuid.UUID       `bun:"type:uuid" json:"specializationId"`

	DoctorID uuid.NullUUID `bun:"type:uuid" json:"doctorId"`
	Doctor   *Doctor       `bun:"rel:belongs-to" json:"doctor"`

	OMS           bool   `json:"oms"`
	Mrt           bool   `json:"mrt"`
	MrtZone       string `json:"mrtZone"`
	MrtAnesthesia bool   `json:"mrtAnesthesia"`

	User    *User     `bun:"rel:belongs-to" json:"user"`
	UserID  uuid.UUID `bun:"type:uuid" json:"userId"`
	Child   *Child    `bun:"rel:belongs-to" json:"child"`
	ChildID uuid.UUID `bun:"type:uuid" json:"childId"`
}

type Appointments []*ApplicationCar

func (item *Appointment) SetForeignKeys() {
	if item.User != nil {
		item.UserID = item.User.ID
	}
	if item.Child != nil {
		item.ChildID = item.Child.ID
	}
}
