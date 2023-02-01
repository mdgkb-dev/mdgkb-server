package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Teacher struct {
	bun.BaseModel `bun:"teachers,select:teachers_view,alias:teachers_view"`
	ID            uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	EmployeeID    uuid.NullUUID `bun:"type:uuid" json:"employeeId,omitempty"`
	Employee      *Employee     `bun:"rel:belongs-to" json:"employee"`
	Position      string        `json:"position"`
	DpoCourses    DpoCourses    `bun:"rel:has-many" json:"dpoCourses"`

	FullName  string `bun:"-" json:"fullName"`
	DateBirth string `bun:"-" json:"dateBirth"`
	IsMale    string `bun:"-" json:"isMale"`
}

type Teachers []*Teacher

func (item *Teacher) SetFilePath(fileID *string) *string {
	path := item.Employee.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}

func (item *Teacher) SetForeignKeys() {
	if item.Employee != nil {
		item.EmployeeID = item.Employee.ID
	}
}
