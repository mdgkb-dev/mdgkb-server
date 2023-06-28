package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ResidencyApplication struct {
	bun.BaseModel `bun:"residency_applications,select:residency_applications_view,alias:residency_applications_view"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	PointsAchievements         int    `json:"pointsAchievements"`
	PointsEntrance             int    `json:"pointsEntrance"`
	ApplicationNum             string `json:"applicationNum"`
	PrimaryAccreditation       bool   `json:"primaryAccreditation"`
	PrimaryAccreditationPoints int    `json:"primaryAccreditationPoints"`
	PrimaryAccreditationPlace  string `json:"primaryAccreditationPlace"`

	MdgkbExam                  bool   `json:"mdgkbExam"`
	EntranceExamPlace          string `json:"entranceExamPlace"`
	EntranceExamSpecialisation string `json:"entranceExamSpecialisation"`

	Main bool `json:"main"`
	Paid bool `json:"paid"`

	UserEdit bool `json:"userEdit"`

	AdmissionCommittee bool `json:"admissionCommittee"`

	Diploma   *Diploma      `bun:"rel:belongs-to" json:"diploma"`
	DiplomaID uuid.NullUUID `bun:"type:uuid" json:"diplomaId"`

	ResidencyCourse   *ResidencyCourse `bun:"rel:belongs-to" json:"residencyCourse"`
	ResidencyCourseID uuid.NullUUID    `bun:"type:uuid" json:"residencyCourseId"`

	ResidencyApplicationPointsAchievements          ResidencyApplicationPointsAchievements `bun:"rel:has-many" json:"residencyApplicationPointsAchievements"`
	ResidencyApplicationPointsAchievementsForDelete []uuid.UUID                            `bun:"-" json:"residencyApplicationPointsAchievementsForDelete"`

	FormValue   *FormValue    `bun:"rel:belongs-to" json:"formValue"`
	FormValueID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"formValueId"`
}

type ResidencyApplications []*ResidencyApplication

type ResidencyApplicationsWithCount struct {
	ResidencyApplications ResidencyApplications `json:"items"`
	Count                 int                   `json:"count"`
}

func (item *ResidencyApplication) SetForeignKeys() {
	item.ResidencyCourseID = item.ResidencyCourse.ID
	item.FormValueID = item.FormValue.ID
	item.DiplomaID = item.Diploma.ID
}

func (item *ResidencyApplication) SetFilePath(fileID *string) *string {
	if item.FormValue != nil {
		path := item.FormValue.SetFilePath(fileID)
		if path != nil {
			return path
		}
	}
	path := item.ResidencyApplicationPointsAchievements.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}

func (item *ResidencyApplication) SetIDForChildren() {
	for i := range item.ResidencyApplicationPointsAchievements {
		item.ResidencyApplicationPointsAchievements[i].ResidencyApplicationID = item.ID
	}
}

func (item *ResidencyApplication) GetCourseName() string {
	name := ""
	if item.ResidencyCourse == nil {
		return name
	}
	for _, course := range item.ResidencyCourse.ResidencyCoursesSpecializations {
		if course.Main {
			name = course.Specialization.Name
		}
	}
	return name
}
