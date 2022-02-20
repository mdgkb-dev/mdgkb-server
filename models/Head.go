package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Head struct {
	bun.BaseModel `bun:"heads,select:heads,alias:heads"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Human         *Human        `bun:"rel:belongs-to" json:"human"`
	HumanId       uuid.NullUUID `bun:"type:uuid" json:"humanId"`
	Position      string        `json:"position"`
	Tags          string        `json:"tags"`
	Photo         *FileInfo     `bun:"rel:belongs-to" json:"photo"`
	PhotoId       uuid.UUID     `bun:"type:uuid" json:"photoId"`

	AcademicDegree    string      `json:"academicDegree"`
	AcademicRank      string      `json:"academicRank"`
	Regalias          Regalias    `bun:"rel:has-many" json:"regalias"`
	RegaliasForDelete []uuid.UUID `bun:"-" json:"regaliasForDelete"`
	Timetable         *Timetable  `bun:"rel:belongs-to" json:"timetable"`

	TimetableId uuid.UUID `bun:"type:uuid" json:"timetableId"`
	IsMain      bool      `json:"isMain"`

	Departments          Departments `bun:"rel:has-many" json:"departments"`
	DepartmentsForDelete []uuid.UUID `bun:"-" json:"departmentsForDelete"`

	ContactInfo   *ContactInfo `bun:"rel:belongs-to" json:"contactInfo"`
	ContactInfoID uuid.UUID    `bun:"type:uuid" json:"contactInfoId"`
}

type Heads []*Head

func (item *Head) SetForeignKeys() {
	if item.Photo != nil {
		item.PhotoId = item.Photo.ID.UUID
	}
	if item.Human != nil {
		item.HumanId = item.Human.ID
	}
	if item.Timetable != nil {
		item.TimetableId = item.Timetable.ID
	}
	if item.ContactInfo != nil {
		item.ContactInfoID = item.ContactInfo.ID
	}
}

func (item *Head) SetIdForChildren() {
	for i := range item.Regalias {
		item.Regalias[i].HeadID = item.ID
	}
	for i := range item.Departments {
		item.Departments[i].HeadID = item.ID
	}
}
