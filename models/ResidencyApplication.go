package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ResidencyApplication struct {
	bun.BaseModel `bun:"residency_applications,select:residency_applications_view,alias:residency_applications_view"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	PointsAchievements int `json:"pointsAchievements"`
	PointsEntrance     int `json:"pointsEntrance"`

	PrimaryAccreditation       bool   `json:"primaryAccreditation"`
	PrimaryAccreditationPoints int    `json:"primaryAccreditationPoints"`
	PrimaryAccreditationPlace  string `json:"primaryAccreditationPlace"`

	Main bool `json:"main"`
	Paid bool `json:"paid"`

	ResidencyCourse   *ResidencyCourse `bun:"rel:belongs-to" json:"residencyCourse"`
	ResidencyCourseID uuid.NullUUID    `bun:"type:uuid" json:"residencyCourseId"`

	ResidencyApplicationPointsAchievements          ResidencyApplicationPointsAchievements `bun:"rel:has-many" json:"residencyApplicationPointsAchievements"`
	ResidencyApplicationPointsAchievementsForDelete []uuid.UUID                            `bun:"-" json:"residencyApplicationPointsAchievementsForDelete"`

	FormValue   *FormValue    `bun:"rel:belongs-to" json:"formValue"`
	FormValueID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"formValueId"`
}

type ResidencyApplications []*ResidencyApplication

type ResidencyApplicationsWithCount struct {
	ResidencyApplications ResidencyApplications `json:"residencyApplications"`
	Count                 int                   `json:"count"`
}

func (item *ResidencyApplication) SetForeignKeys() {
	item.ResidencyCourseID = item.ResidencyCourse.ID
	item.FormValueID = item.FormValue.ID
}

func (item *ResidencyApplication) SetFilePath(fileID *string) *string {
	path := item.FormValue.SetFilePath(fileID)
	if path != nil {
		return path
	}
	path = item.ResidencyApplicationPointsAchievements.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}

func (item *ResidencyApplication) SetIdForChildren() {
	for i := range item.ResidencyApplicationPointsAchievements {
		item.ResidencyApplicationPointsAchievements[i].ResidencyApplicationID = item.ID
	}
}

func (item *ResidencyApplication) GetCourseName() string {
	name := ""
	for _, course := range item.ResidencyCourse.ResidencyCoursesSpecializations {
		if course.Main {
			fmt.Println(course.Specialization.Name)
			name = course.Specialization.Name
		}
	}
	return name
}
