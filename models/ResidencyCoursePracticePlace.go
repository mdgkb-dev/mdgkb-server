package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ResidencyCoursePracticePlace struct {
	bun.BaseModel `bun:"residency_course_practice_places,select:residency_course_practice_places,alias:residency_course_practice_places"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name          string        `json:"name"`
	Link          string        `json:"link"`
	Order         uint8         `bun:"residency_course_practice_place_order" json:"order"`

	ResidencyCoursePracticePlaceGroup   *ResidencyCoursePracticePlaceGroup `bun:"rel:belongs-to" json:"residencyCoursePracticePlaceGroup"`
	ResidencyCoursePracticePlaceGroupID uuid.NullUUID                      `bun:"type:uuid" json:"residencyCoursePracticePlaceGroupId"`

	Division   *Division     `bun:"rel:belongs-to" json:"division"`
	DivisionID uuid.NullUUID `bun:"type:uuid" json:"divisionId,omitempty"`
}

type ResidencyCoursePracticePlaces []*ResidencyCoursePracticePlace

func (item *ResidencyCoursePracticePlace) SetForeignKeys() {
	if item.Division != nil {
		item.DivisionID = item.Division.ID
	}
}

func (items ResidencyCoursePracticePlaces) SetForeignKeys() {
	for i := range items {
		items[i].SetForeignKeys()
	}
}
