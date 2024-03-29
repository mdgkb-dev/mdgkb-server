package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Head struct {
	bun.BaseModel `bun:"heads,select:heads_view,alias:heads_view"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	Employee   *Employee     `bun:"rel:belongs-to" json:"employee"`
	EmployeeID uuid.NullUUID `bun:"type:uuid" json:"employeeId,omitempty"`

	Position string `json:"position"`
	Tags     string `json:"tags"`

	Timetable   *Timetable    `bun:"rel:belongs-to" json:"timetable"`
	TimetableID uuid.NullUUID `bun:"type:uuid" json:"timetableId"`

	IsMain bool `json:"isMain"`

	Departments          Departments `bun:"rel:has-many" json:"departments"`
	DepartmentsForDelete []uuid.UUID `bun:"-" json:"departmentsForDelete"`

	ContactInfo   *ContactInfo `bun:"rel:belongs-to" json:"contactInfo"`
	ContactInfoID uuid.UUID    `bun:"type:uuid" json:"contactInfoId"`

	FullName  string `bun:"-" json:"fullName"`
	DateBirth string `bun:"-" json:"dateBirth"`
	IsMale    string `bun:"-" json:"isMale"`
	Order     uint   `bun:"item_order" json:"order"`
}

type Heads []*Head

func (item *Head) SetForeignKeys() {
	if item.Timetable != nil {
		item.TimetableID = item.Timetable.ID
	}
	if item.ContactInfo != nil {
		item.ContactInfoID = item.ContactInfo.ID
	}
}

func (item *Head) SetIDForChildren() {
	for i := range item.Departments {
		item.Departments[i].HeadID = item.ID
	}
}
