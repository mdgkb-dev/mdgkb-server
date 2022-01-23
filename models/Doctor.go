package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/helpers/uploadHelper"
)

type Doctor struct {
	bun.BaseModel     `bun:"doctors,select:doctors_view,alias:doctors_view"`
	ID                uuid.NullUUID   `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Division          *Division       `bun:"rel:belongs-to" json:"division"`
	DivisionId        uuid.NullUUID   `bun:"type:uuid" json:"divisionId,omitempty"`
	Human             *Human          `bun:"rel:belongs-to" json:"human"`
	HumanId           uuid.UUID       `bun:"type:uuid" json:"humanId"`
	Position          *Position       `bun:"rel:belongs-to" json:"position"`
	PositionID        uuid.UUID       `bun:"type:uuid" json:"positionId"`
	Schedule          string          `json:"schedule"`
	Tags              string          `json:"tags"`
	MedicalProfile    *MedicalProfile `bun:"rel:belongs-to" json:"medicalProfile"`
	MedicalProfileID  uuid.UUID       `bun:"type:uuid" json:"medicalProfileId"`
	Order             int             `bun:"item_order" json:"order"'`
	FileInfo          *FileInfo       `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoId        uuid.UUID       `bun:"type:uuid" json:"fileInfoId"`
	PhotoMini         *FileInfo       `bun:"rel:belongs-to" json:"photoMini"`
	PhotoMiniID       uuid.UUID       `bun:"type:uuid" json:"photoMiniId"`
	DoctorComments    DoctorComments  `bun:"rel:has-many" json:"doctorComments"`
	NewsDoctors       NewsDoctors     `bun:"rel:has-many" json:"newsDoctors"`
	MosDoctorLink     string          `json:"mosDoctorLink"`
	OnlineDoctorID    string          `json:"onlineDoctorId"`
	AcademicDegree    string          `json:"academicDegree"`
	AcademicRank      string          `json:"academicRank"`
	RegaliasCount     int             `bun:"-" json:"regaliasCount"`
	CommentsCount     int             `bun:"-" json:"commentsCount"`
	Show              bool            `json:"show"`
	Regalias          Regalias        `bun:"rel:has-many" json:"regalias"`
	RegaliasForDelete []uuid.UUID     `bun:"-" json:"regaliasForDelete"`

	Educations          Educations  `bun:"rel:has-many" json:"educations"`
	EducationsForDelete []uuid.UUID `bun:"-" json:"educationsForDelete"`

	Experiences          Experiences `bun:"rel:has-many" json:"experiences"`
	ExperiencesForDelete []uuid.UUID `bun:"-" json:"experiencesForDelete"`

	Certificates          Certificates `bun:"rel:has-many" json:"certificates"`
	CertificatesForDelete []uuid.UUID  `bun:"-" json:"certificatesForDelete"`

	DoctorPaidServices          DoctorPaidServices `bun:"rel:has-many" json:"doctorPaidServices"`
	DoctorPaidServicesForDelete []uuid.UUID        `bun:"-" json:"doctorPaidServicesForDelete"`

	Timetable   *Timetable `bun:"rel:belongs-to" json:"timetable"`
	TimetableId uuid.UUID  `bun:"type:uuid,nullzero,default:NULL" json:"timetableId"`
}

type Doctors []*Doctor

func (item *Doctor) SetFilePath(fileID *string) *string {
	for i := range item.Certificates {
		if item.Certificates[i].Scan.ID.UUID.String() == *fileID {
			item.Certificates[i].Scan.FileSystemPath = uploadHelper.BuildPath(fileID)
			return &item.Certificates[i].Scan.FileSystemPath
		}
	}
	if item.FileInfo.ID.UUID.String() == *fileID {
		item.FileInfo.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.FileInfo.FileSystemPath
	}
	if item.PhotoMini.ID.UUID.String() == *fileID {
		item.PhotoMini.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.PhotoMini.FileSystemPath
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
	if item.PhotoMini != nil {
		item.PhotoMiniID = item.PhotoMini.ID.UUID
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
	for i := range item.DoctorPaidServices {
		item.DoctorPaidServices[i].DoctorID = item.ID
	}
}
