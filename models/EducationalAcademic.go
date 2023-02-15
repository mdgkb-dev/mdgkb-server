package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type EducationalAcademic struct {
	bun.BaseModel `bun:"educational_academics,select:educational_academics_view,alias:educational_academics_view"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	EmployeeID    uuid.NullUUID `bun:"type:uuid" json:"employeeId"`
	Employee      *Employee     `bun:"rel:belongs-to" json:"employee"`
	Order         int           `bun:"item_order" json:"order"`
	FullName      string        `bun:"-" json:"fullName"`
}

type EducationalAcademics []*EducationalAcademic
