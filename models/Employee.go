package models

import (
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"github.com/uptrace/bun"
)

type Employee struct {
	bun.BaseModel `bun:"employees,select:employees_view,alias:employees_view"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	Human   *Human        `bun:"rel:belongs-to" json:"human"`
	HumanID uuid.NullUUID `bun:"type:uuid" json:"humanId"`

	AcademicDegree string `json:"academicDegree"`
	AcademicRank   string `json:"academicRank"`

	PartTime bool `json:"partTime"`

	Regalias          Regalias    `bun:"rel:has-many" json:"regalias"`
	RegaliasForDelete []uuid.UUID `bun:"-" json:"regaliasForDelete"`

	Educations          Educations  `bun:"rel:has-many" json:"educations"`
	EducationsForDelete []uuid.UUID `bun:"-" json:"educationsForDelete"`

	Experiences          Experiences `bun:"rel:has-many" json:"experiences"`
	ExperiencesForDelete []uuid.UUID `bun:"-" json:"experiencesForDelete"`

	Certificates          Certificates `bun:"rel:has-many" json:"certificates"`
	CertificatesForDelete []uuid.UUID  `bun:"-" json:"certificatesForDelete"`

	TeachingActivities          TeachingActivities `bun:"rel:has-many" json:"teachingActivities"`
	TeachingActivitiesForDelete []uuid.UUID        `bun:"-" json:"teachingActivitiesForDelete"`

	Certifications          Certifications `bun:"rel:has-many" json:"certifications"`
	CertificationsForDelete []uuid.UUID    `bun:"-" json:"certificationsForDelete"`

	Accreditations          Accreditations `bun:"rel:has-many" json:"accreditations"`
	AccreditationsForDelete []uuid.UUID    `bun:"-" json:"accreditationsForDelete"`

	FullName  string `bun:"-" json:"fullName"`
	IsMale    string `bun:"-" json:"isMale"`
	DateBirth string `bun:"-" json:"dateBirth"`
}

func (item *Employee) SetFilePath(fileID *string) *string {
	for i := range item.Certificates {
		if item.Certificates[i].Scan.ID.UUID.String() == *fileID {
			item.Certificates[i].Scan.FileSystemPath = uploadHelper.BuildPath(fileID)
			return &item.Certificates[i].Scan.FileSystemPath
		}
	}
	path := item.Human.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}

type Employees []*Employee

type EmployeesWithCount struct {
	Employees Employees `json:"items"`
	Count     int       `json:"count"`
}

func (item *Employee) SetForeignKeys() {
	if item.Human != nil {
		item.HumanID = item.Human.ID
	}
}

func (item *Employee) SetIDForChildren() {
	for i := range item.Educations {
		item.Educations[i].EmployeeID = item.ID
	}
	for i := range item.Experiences {
		item.Experiences[i].EmployeeID = item.ID
	}
	for i := range item.Certificates {
		item.Certificates[i].EmployeeID = item.ID
	}
	for i := range item.Regalias {
		item.Regalias[i].EmployeeID = item.ID
	}

	for i := range item.TeachingActivities {
		item.TeachingActivities[i].EmployeeID = item.ID
	}

	for i := range item.TeachingActivities {
		item.TeachingActivities[i].EmployeeID = item.ID
	}

	for i := range item.Certifications {
		item.Certifications[i].EmployeeID = item.ID
	}

	for i := range item.Accreditations {
		item.Accreditations[i].EmployeeID = item.ID
	}
}
