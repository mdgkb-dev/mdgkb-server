package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ResidencyApplication struct {
	bun.BaseModel `bun:"residency_applications,select:residency_applications_view,alias:residency_applications_view"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	PointsAchievements int `json:"pointsAchievements"`
	PointsEntrance     int `json:"pointsEntrance"`

	ResidencyCourse   *ResidencyCourse `bun:"rel:belongs-to" json:"residencyCourse"`
	ResidencyCourseID uuid.NullUUID    `bun:"type:uuid" json:"residencyCourseId"`

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
	return path
}

func (item *ResidencyApplication) SetIdForChildren() {
}
