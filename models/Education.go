package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Education struct {
	bun.BaseModel `bun:"educations,alias:educations"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Type          string    `json:"type"`
	Institution   string    `json:"institution"`
	Document      string    `json:"document"`

	EducationSpeciality   *EducationSpeciality `bun:"rel:belongs-to" json:"educationSpeciality"`
	EducationSpecialityID uuid.UUID            `bun:"type:uuid" json:"educationSpecialityId"`

	EducationCertification   *EducationCertification `bun:"rel:belongs-to" json:"educationCertification"`
	EducationCertificationID uuid.UUID               `bun:"type:uuid" json:"educationCertificationId"`

	EducationQualificationID uuid.UUID               `bun:"type:uuid" json:"educationQualificationId"`
	EducationQualification   *EducationQualification `bun:"rel:belongs-to" json:"educationQualification"`

	DoctorID uuid.UUID `bun:"type:uuid" json:"doctorId"`
	Doctor   *Doctor   `bun:"rel:belongs-to" json:"doctor"`
}

type Educations []*Education

func (item *Education) SetIdForChildren() {
	if item.EducationCertification != nil {
		item.EducationCertificationID = item.EducationCertification.ID
	}
	if item.EducationQualification != nil {
		item.EducationQualificationID = item.EducationQualification.ID
	}
}

func (items Educations) SetIdForChildren() {
	for i := range items {
		items[i].SetIdForChildren()
	}
}

func (items Educations) GetEducationCertifications() EducationCertifications {
	itemsForGet := make(EducationCertifications, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.EducationCertification)
	}
	return itemsForGet
}

func (items Educations) GetEducationQualification() EducationQualifications {
	itemsForGet := make(EducationQualifications, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.EducationQualification)
	}
	return itemsForGet
}
