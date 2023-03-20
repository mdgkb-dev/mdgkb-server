package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ResidencyCoursePracticePlaceGroup struct {
	bun.BaseModel `bun:"residency_course_practice_place_groups,select:residency_course_practice_place_groups,alias:residency_course_practice_place_groups"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name          string        `json:"name"`
	Link          string        `json:"link"`
	Order         uint8         `bun:"item_order" json:"order"`

	ResidencyCourse   *ResidencyCourse `bun:"rel:belongs-to" json:"residencyCourse"`
	ResidencyCourseID uuid.NullUUID    `bun:"type:uuid" json:"residencyCourseId"`

	ResidencyCoursePracticePlaces          ResidencyCoursePracticePlaces `bun:"rel:has-many" json:"residencyCoursePracticePlaces"`
	ResidencyCoursePracticePlacesForDelete []uuid.UUID                   `bun:"-" json:"residencyCoursePracticePlacesForDelete"`
}

type ResidencyCoursePracticePlaceGroups []*ResidencyCoursePracticePlaceGroup

func (item *ResidencyCoursePracticePlaceGroup) SetForeignKeys() {
	if item.ResidencyCourse != nil {
		item.ResidencyCourseID = item.ResidencyCourse.ID
	}
}

func (items ResidencyCoursePracticePlaceGroups) SetForeignKeys() {
	for i := range items {
		items[i].SetForeignKeys()
	}
}

func (item *ResidencyCoursePracticePlaceGroup) SetIDForChildren() {
	for i := range item.ResidencyCoursePracticePlaces {
		item.ResidencyCoursePracticePlaces[i].ResidencyCoursePracticePlaceGroupID = item.ID
	}
}

func (items ResidencyCoursePracticePlaceGroups) SetIDForChildren() {
	for i := range items {
		items[i].SetIDForChildren()
	}
}

func (items ResidencyCoursePracticePlaceGroups) GetResidencyCoursePracticePlaces() ResidencyCoursePracticePlaces {
	itemsForGet := make(ResidencyCoursePracticePlaces, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].ResidencyCoursePracticePlaces...)
	}
	return itemsForGet
}

func (items ResidencyCoursePracticePlaceGroups) GetResidencyCoursePracticePlacesForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.ResidencyCoursePracticePlacesForDelete...)
	}
	return itemsForGet
}
