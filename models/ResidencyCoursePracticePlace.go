package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ResidencyCoursePracticePlace struct {
	bun.BaseModel     `bun:"residency_course_practice_places,select:residency_course_practice_places,alias:residency_course_practice_places"`
	ID                uuid.NullUUID    `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name              string           `json:"name"`
	Link              string           `json:"link"`
	Order             string           `bun:"residency_course_practice_place_order" json:"order"`
	ResidencyCourse   *ResidencyCourse `bun:"rel:belongs-to" json:"residencyCourse"`
	ResidencyCourseID uuid.NullUUID    `bun:"type:uuid" json:"residencyCourseId"`
}

type ResidencyCoursePracticePlaces []*ResidencyCoursePracticePlace
