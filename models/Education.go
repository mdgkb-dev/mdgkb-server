package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Education struct {
	bun.BaseModel         `bun:"educations,alias:educations"`
	ID                    uuid.UUID            `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Type                  string               `json:"type"`
	Institution           string               `json:"institution"`
	Document              string               `json:"document"`
	Qualification         string               `json:"qualification"`
	Specialization        string               `json:"specialization"`
	Start                 time.Time            `bun:"education_start" json:"start"`
	End                   time.Time            `bun:"education_end" json:"end"`
	EducationSpeciality   *EducationSpeciality `bun:"rel:belongs-to" json:"educationSpeciality"`
	EducationSpecialityID uuid.UUID            `bun:"type:uuid" json:"educationSpecialityId"`

	EducationCertification   *EducationCertification `bun:"rel:belongs-to" json:"educationCertification"`
	EducationCertificationID uuid.NullUUID           `bun:"type:uuid" json:"educationCertificationId"`

	EducationAccreditationID uuid.NullUUID           `bun:"type:uuid" json:"educationAccreditationId"`
	EducationAccreditation   *EducationAccreditation `bun:"rel:belongs-to" json:"educationAccreditation"`

	EmployeeID uuid.NullUUID `bun:"type:uuid" json:"employeeId"`
	Employee   *Employee     `bun:"rel:belongs-to" json:"employee"`
}

type Educations []*Education

func (item *Education) SetForeignKeys() {
	if item.EducationCertification != nil {
		item.EducationCertificationID = item.EducationCertification.ID
	}
	if item.EducationAccreditation != nil {
		item.EducationAccreditationID = item.EducationAccreditation.ID
	}
}

func (items Educations) SetForeignKeys() {
	for i := range items {
		items[i].SetForeignKeys()
	}
}

func (items Educations) GetEducationCertifications() EducationCertifications {
	itemsForGet := make(EducationCertifications, 0)
	for _, item := range items {
		if item.EducationCertification != nil {
			itemsForGet = append(itemsForGet, item.EducationCertification)
		}
	}
	return itemsForGet
}

func (items Educations) GetEducationQualification() EducationAccreditations {
	itemsForGet := make(EducationAccreditations, 0)
	for _, item := range items {
		if item.EducationAccreditation != nil {
			itemsForGet = append(itemsForGet, item.EducationAccreditation)
		}
	}
	return itemsForGet
}
