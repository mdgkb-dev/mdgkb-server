package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Doctor struct {
	bun.BaseModel  `bun:"doctors,select:doctors_view,alias:doctors_view"`
	ID             uuid.UUID      `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Division       *Division      `bun:"rel:belongs-to" json:"division"`
	DivisionId     uuid.UUID      `bun:"type:uuid,nullzero,default:NULL" json:"divisionId,omitempty"`
	Human          *Human         `bun:"rel:belongs-to" json:"human"`
	HumanId        uuid.UUID      `bun:"type:uuid" json:"humanId"`
	Schedule       string         `json:"schedule"`
	Position       string         `json:"position"`
	Tags           string         `json:"tags"`
	FileInfo       *FileInfo      `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoId     uuid.UUID      `bun:"type:uuid" json:"fileInfoId"`
	DoctorComments DoctorComments `bun:"rel:has-many" json:"doctorComments"`

	AcademicDegree          string         `json:"academicDegree"`
	AcademicRank            string         `json:"academicRank"`
	DoctorRegalias          DoctorRegalias `bun:"rel:has-many" json:"doctorRegalias"`
	DoctorRegaliasForDelete []uuid.UUID    `bun:"-" json:"doctorRegaliasForDelete"`

	Educations          Educations  `bun:"rel:has-many" json:"educations"`
	EducationsForDelete []uuid.UUID `bun:"-" json:"educationsForDelete"`

	Timetable   *Timetable `bun:"rel:belongs-to" json:"timetable"`
	TimetableId uuid.UUID  `bun:"type:uuid,nullzero,default:NULL" json:"timetableId"`
}

type Doctors []*Doctor

func (item *Doctor) SetForeignKeys() {
	if item.FileInfo != nil {
		item.FileInfoId = item.FileInfo.ID.UUID
	}
	if item.Human != nil {
		item.HumanId = item.Human.ID
	}
	if item.Timetable != nil {
		item.TimetableId = item.Timetable.ID
	}
}

func (item *Doctor) SetIdForChildren() {
	for i := range item.Educations {
		item.Educations[i].DoctorID = item.ID
	}
	for i := range item.DoctorRegalias {
		item.DoctorRegalias[i].DoctorID = item.ID
	}
}
