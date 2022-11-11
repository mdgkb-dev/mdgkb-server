package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Doctor struct {
	bun.BaseModel    `bun:"doctors,select:doctors_view,alias:doctors_view"`
	ID               uuid.NullUUID   `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Description      string          `json:"description"`
	Specialization   *Specialization `bun:"rel:belongs-to" json:"specialization"`
	SpecializationID uuid.NullUUID   `bun:"type:uuid" json:"specializationId,omitempty"`

	Employee   *Employee     `bun:"rel:belongs-to" json:"employee"`
	EmployeeID uuid.NullUUID `bun:"type:uuid" json:"employeeId,omitempty"`

	Position         *Position       `bun:"rel:belongs-to" json:"position"`
	PositionID       uuid.NullUUID   `bun:"type:uuid" json:"positionId"`
	Schedule         string          `json:"schedule"`
	Tags             string          `json:"tags"`
	MedicalProfile   *MedicalProfile `bun:"rel:belongs-to" json:"medicalProfile"`
	MedicalProfileID uuid.NullUUID   `bun:"type:uuid" json:"medicalProfileId"`
	Order            int             `bun:"item_order" json:"order"`

	DoctorComments            DoctorComments   `bun:"rel:has-many" json:"doctorComments"`
	NewsDoctors               NewsDoctors      `bun:"rel:has-many" json:"newsDoctors"`
	DoctorsDivisions          DoctorsDivisions `bun:"rel:has-many" json:"doctorsDivisions"`
	DoctorsDivisionsForDelete []uuid.UUID      `bun:"-" json:"doctorsDivisionsForDelete"`
	MosDoctorLink             string           `json:"mosDoctorLink"`
	OnlineDoctorID            string           `json:"onlineDoctorId"`

	DoctorPaidServices          DoctorPaidServices `bun:"rel:has-many" json:"doctorPaidServices"`
	DoctorPaidServicesForDelete []uuid.UUID        `bun:"-" json:"doctorPaidServicesForDelete"`
	HasAppointment              bool               `json:"hasAppointment"`
	Timetable                   *Timetable         `bun:"rel:belongs-to" json:"timetable"`
	TimetableID                 uuid.NullUUID      `bun:"type:uuid,nullzero,default:NULL" json:"timetableId"`
}

type Doctors []*Doctor

func (item *Doctor) SetFilePath(fileID *string) *string {
	path := item.Employee.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}

func (item *Doctor) SetForeignKeys() {
	if item.Position != nil {
		item.PositionID = item.Position.ID
	}
	if item.Timetable != nil {
		item.TimetableID = item.Timetable.ID
	}
	if item.MedicalProfile != nil {
		item.MedicalProfileID = item.MedicalProfile.ID
	}
	if item.Employee != nil {
		item.EmployeeID = item.Employee.ID
	}
}

func (item *Doctor) SetIDForChildren() {
	for i := range item.DoctorPaidServices {
		item.DoctorPaidServices[i].DoctorID = item.ID
	}
	for i := range item.DoctorsDivisions {
		item.DoctorsDivisions[i].DoctorID = item.ID
	}
}

func (item *Doctor) InitAppointmentsSlots() {
	if item.Timetable != nil {
		item.Timetable.InitAppointmentsSlots()
	}
}

func (items Doctors) InitAppointmentsSlots() {
	for i := range items {
		items[i].InitAppointmentsSlots()
	}
}

func (item *Doctor) InitAppointments(days []time.Time) Appointments {
	appointments := make(Appointments, 0)
	if item.Timetable == nil {
		return appointments
	}
	//for _, day := range days {
	//	for _, weekday := range item.Timetable.TimetableDays {
	//		for _, slot := range weekday.AppointmentsSlots {
	//			apointment := Appointment{}
	//timeToAppointment, _ := time.Parse("2006-01-02 15:04", day.Format("2006-01-02")+" "+slot)
	//apointment.Time = timeToAppointment
	//apointment.DoctorID = item.ID
	//appointments = append(appointments, &apointment)
	//}
	//}
	//}
	return appointments
}

func (items Doctors) InitAppointments(days []time.Time) Appointments {
	appointments := make(Appointments, 0)
	for i := range items {
		appointments = append(appointments, items[i].InitAppointments(days)...)
	}
	return appointments
}
