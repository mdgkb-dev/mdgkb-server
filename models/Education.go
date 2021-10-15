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
