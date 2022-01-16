package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/helpers/uploadHelper"
)

type Doctor struct {
	bun.BaseModel     `bun:"doctors,select:doctors_view,alias:doctors_view"`
	ID                uuid.NullUUID  `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Division          *Division      `bun:"rel:belongs-to" json:"division"`
	DivisionId        uuid.NullUUID  `bun:"type:uuid" json:"divisionId,omitempty"`
	Human             *Human         `bun:"rel:belongs-to" json:"human"`
	HumanId           uuid.UUID      `bun:"type:uuid" json:"humanId"`
	Schedule          string         `json:"schedule"`
	Position          string         `json:"position"`
	Tags              string         `json:"tags"`
	FileInfo          *FileInfo      `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoId        uuid.UUID      `bun:"type:uuid" json:"fileInfoId"`
	DoctorComments    DoctorComments `bun:"rel:has-many" json:"doctorComments"`
	NewsDoctors       NewsDoctors    `bun:"rel:has-many" json:"newsDoctors"`
	MosDoctorLink     string         `json:"mosDoctorLink"`
	AcademicDegree    string         `json:"academicDegree"`
	AcademicRank      string         `json:"academicRank"`
	Show              bool           `json:"show"`
	Regalias          Regalias       `bun:"rel:has-many" json:"regalias"`
	RegaliasForDelete []uuid.UUID    `bun:"-" json:"regaliasForDelete"`

	Educations          Educations  `bun:"rel:has-many" json:"educations"`
	EducationsForDelete []uuid.UUID `bun:"-" json:"educationsForDelete"`

	Experiences          Experiences `bun:"rel:has-many" json:"experiences"`
	ExperiencesForDelete []uuid.UUID `bun:"-" json:"experiencesForDelete"`

	Certificates          Certificates `bun:"rel:has-many" json:"certificates"`
	CertificatesForDelete []uuid.UUID  `bun:"-" json:"certificatesForDelete"`

	Timetable   *Timetable `bun:"rel:belongs-to" json:"timetable"`
	TimetableId uuid.UUID  `bun:"type:uuid,nullzero,default:NULL" json:"timetableId"`
}

type Doctors []*Doctor

func (item *Doctor) SetFilePath(fileID *string) *string {
	fmt.Println(item.Certificates)
	fmt.Println(item.Certificates)
	fmt.Println(item.Certificates)
	fmt.Println(item.Certificates)
	for i := range item.Certificates {
		fmt.Println(item.Certificates[i].Scan.ID, *fileID)
		fmt.Println(item.Certificates[i].Scan.ID, *fileID)
		fmt.Println(item.Certificates[i].Scan.ID, *fileID)
		fmt.Println(item.Certificates[i].Scan.ID, *fileID)
		fmt.Println(item.Certificates[i].Scan.ID, *fileID)
		fmt.Println(item.Certificates[i].Scan.ID, *fileID)
		fmt.Println(item.Certificates[i].Scan.ID, *fileID)
		if item.Certificates[i].Scan.ID.UUID.String() == *fileID {
			item.Certificates[i].Scan.FileSystemPath = uploadHelper.BuildPath(fileID)
			return &item.Certificates[i].Scan.FileSystemPath
		}
	}
	if item.FileInfo.ID.UUID.String() == *fileID {
		item.FileInfo.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.FileInfo.FileSystemPath
	}
	return nil
}

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
	for i := range item.Experiences {
		item.Experiences[i].DoctorID = item.ID
	}
	for i := range item.Certificates {
		item.Certificates[i].DoctorID = item.ID
	}
	for i := range item.Regalias {
		item.Regalias[i].DoctorID = item.ID
	}
}
