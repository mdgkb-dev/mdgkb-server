package models

import "github.com/google/uuid"

type Doctor struct {
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
	DoctorRegalias          DoctorRegalias `bun:"type:has-many" json:"doctorRegalias"`
	DoctorRegaliasForDelete []uuid.UUID    `bun:"type:has-many" json:"doctorRegaliasForDelete"`

	Educations          Educations  `bun:"type:has-many" json:"educations"`
	EducationsForDelete []uuid.UUID `bun:"type:has-many" json:"educationsForDelete"`
}

type Doctors []*Doctor

func (item *Doctor) SetForeignKeys() {
	if item.FileInfo != nil {
		item.FileInfoId = item.FileInfo.ID.UUID
	}
}

func (item *Doctor) SetIdForChildren() {
	for i := range item.Educations {
		item.Educations[i].DoctorID = item.ID
	}
}
