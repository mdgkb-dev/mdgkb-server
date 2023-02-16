package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type NmoCourse struct {
	bun.BaseModel                      `bun:"nmo_courses,select:nmo_courses,alias:nmo_courses"`
	ID                                 uuid.NullUUID             `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Slug                               string                    `json:"slug"`
	Name                               string                    `json:"name"`
	Description                        string                    `json:"description"`
	Order                              int                       `bun:"dpo_course_order" json:"order"`
	Link                               string                    `json:"link"`
	Listeners                          int                       `json:"listeners"`
	Hours                              int                       `json:"hours"`
	Cost                               int                       `json:"cost"`
	MinStart                           time.Time                 `bun:"min_dpo_course_start,scanonly" json:"minStart"`
	Specialization                     *Specialization           `bun:"rel:belongs-to" json:"specialization"`
	SpecializationID                   uuid.NullUUID             `bun:"type:uuid" json:"specializationId,omitempty"`
	NmoCoursesSpecializations          NmoCoursesSpecializations `bun:"rel:has-many" json:"nmoCoursesSpecializations"`
	NmoCoursesSpecializationsForDelete []uuid.UUID               `bun:"-" json:"nmoCoursesSpecializationsForDelete"`
	NmoCoursesTeachers                 NmoCoursesTeachers        `bun:"rel:has-many" json:"nmoCoursesTeachers"`
	NmoCoursesTeachersForDelete        []uuid.UUID               `bun:"-" json:"nmoCoursesTeachersForDelete"`
	// Временно не нужны, оставили на будущие заявки
	//NmoCoursesDates                    NmoCoursesDates           `bun:"rel:has-many" json:"nmoCoursesDates"`
	//NmoCoursesDatesForDelete           []uuid.UUID               `bun:"-" json:"nmoCoursesDatesForDelete"`
	FormPattern   *FormPattern  `bun:"rel:belongs-to" json:"formPattern"`
	FormPatternID uuid.NullUUID `bun:"type:uuid" json:"formPatternId"`
}

type NmoCourses []*NmoCourse

func (item *NmoCourse) SetIDForChildren() {
	for i := range item.NmoCoursesTeachers {
		item.NmoCoursesTeachers[i].NmoCourseID = item.ID
	}
	for i := range item.NmoCoursesSpecializations {
		item.NmoCoursesSpecializations[i].NmoCourseID = item.ID
	}
	//for i := range item.NmoCoursesDates {
	//	item.NmoCoursesDates[i].NmoCourseID = item.ID
	//}
}

func (item *NmoCourse) SetForeignKeys() {
	item.FormPatternID = item.FormPattern.ID
}

func (item *NmoCourse) SetFilePath(fileID *string) *string {
	return item.FormPattern.SetFilePath(fileID)
}
