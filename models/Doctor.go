package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"github.com/uptrace/bun"
)

type Doctor struct {
	bun.BaseModel `bun:"doctors,select:doctors_view,alias:doctors_view"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Division      *Division     `bun:"rel:belongs-to" json:"division"`
	DivisionID    uuid.NullUUID `bun:"type:uuid" json:"divisionId,omitempty"`

	Specialization   *Specialization `bun:"rel:belongs-to" json:"specialization"`
	SpecializationID uuid.NullUUID   `bun:"type:uuid" json:"specializationId,omitempty"`

	Human            *Human          `bun:"rel:belongs-to" json:"human"`
	HumanID          uuid.NullUUID   `bun:"type:uuid" json:"humanId"`
	Position         *Position       `bun:"rel:belongs-to" json:"position"`
	PositionID       uuid.NullUUID   `bun:"type:uuid" json:"positionId"`
	Schedule         string          `json:"schedule"`
	Tags             string          `json:"tags"`
	MedicalProfile   *MedicalProfile `bun:"rel:belongs-to" json:"medicalProfile"`
	MedicalProfileID uuid.NullUUID   `bun:"type:uuid" json:"medicalProfileId"`
	Order            int             `bun:"item_order" json:"order"`

	DoctorComments    DoctorComments `bun:"rel:has-many" json:"doctorComments"`
	NewsDoctors       NewsDoctors    `bun:"rel:has-many" json:"newsDoctors"`
	MosDoctorLink     string         `json:"mosDoctorLink"`
	OnlineDoctorID    string         `json:"onlineDoctorId"`
	AcademicDegree    string         `json:"academicDegree"`
	AcademicRank      string         `json:"academicRank"`
	RegaliasCount     int            `bun:"-" json:"regaliasCount"`
	CommentsCount     int            `bun:"-" json:"commentsCount"`
	Show              bool           `json:"show"`
	Regalias          Regalias       `bun:"rel:has-many" json:"regalias"`
	RegaliasForDelete []uuid.UUID    `bun:"-" json:"regaliasForDelete"`

	Educations          Educations  `bun:"rel:has-many" json:"educations"`
	EducationsForDelete []uuid.UUID `bun:"-" json:"educationsForDelete"`

	Experiences          Experiences `bun:"rel:has-many" json:"experiences"`
	ExperiencesForDelete []uuid.UUID `bun:"-" json:"experiencesForDelete"`

	Certificates          Certificates `bun:"rel:has-many" json:"certificates"`
	CertificatesForDelete []uuid.UUID  `bun:"-" json:"certificatesForDelete"`

	DoctorPaidServices          DoctorPaidServices `bun:"rel:has-many" json:"doctorPaidServices"`
	DoctorPaidServicesForDelete []uuid.UUID        `bun:"-" json:"doctorPaidServicesForDelete"`

	Timetable   *Timetable `bun:"rel:belongs-to" json:"timetable"`
	TimetableID uuid.UUID  `bun:"type:uuid,nullzero,default:NULL" json:"timetableId"`

	EducationalOrganizationAcademic *EducationalOrganizationAcademic `bun:"rel:has-one" json:"educationalOrganizationAcademic"`
}

type Doctors []*Doctor

func (item *Doctor) SetFilePath(fileID *string) *string {
	for i := range item.Certificates {
		if item.Certificates[i].Scan.ID.UUID.String() == *fileID {
			item.Certificates[i].Scan.FileSystemPath = uploadHelper.BuildPath(fileID)
			return &item.Certificates[i].Scan.FileSystemPath
		}
	}
	path := item.Human.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}

func (item *Doctor) SetForeignKeys() {
	if item.Human != nil {
		item.HumanID = item.Human.ID
	}
	if item.Position != nil {
		item.PositionID = item.Position.ID
	}
	if item.Timetable != nil {
		item.TimetableID = item.Timetable.ID
	}
	if item.MedicalProfile != nil {
		item.MedicalProfileID = item.MedicalProfile.ID
	}
}

func (item *Doctor) SetIDForChildren() {
	for i := range item.Educations {
		item.Educations[i].DoctorID = item.ID
	}
	for i := range item.Experiences {
		item.Experiences[i].DoctorID = item.ID
	}
	for i := range item.Certificates {
		item.Certificates[i].DoctorID = item.ID
	}
	for i := range item.Regalias {
		item.Regalias[i].DoctorID = item.ID
	}
	for i := range item.DoctorPaidServices {
		item.DoctorPaidServices[i].DoctorID = item.ID
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
