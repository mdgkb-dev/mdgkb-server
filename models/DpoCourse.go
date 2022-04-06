package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type DpoCourse struct {
	bun.BaseModel                      `bun:"dpo_courses,select:dpo_courses_view,alias:dpo_courses_view"`
	ID                                 uuid.NullUUID             `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	Slug                               string                    `json:"slug"`
	Name                               string                    `json:"name"`
	Description                        string                    `json:"description"`
	Order                              int                       `bun:"dpo_course_order" json:"order"`
	IsNmo                              bool                      `json:"isNmo"`
	LinkNmo                            string                    `json:"linkNmo"`
	Listeners                          int                       `json:"listeners"`
	Hours                              int                       `json:"hours"`
	Cost                               int                       `json:"cost"`
	MinStart                           time.Time                 `bun:"min_dpo_course_start,scanonly" json:"minStart"`
	Specialization                     *Specialization           `bun:"rel:belongs-to" json:"specialization"`
	SpecializationID                   uuid.NullUUID             `bun:"type:uuid" json:"specializationId,omitempty"`
	DpoCoursesSpecializations          DpoCoursesSpecializations `bun:"rel:has-many" json:"dpoCoursesSpecializations"`
	DpoCoursesSpecializationsForDelete []uuid.UUID               `bun:"-" json:"dpoCoursesSpecializationsForDelete"`
	DpoCoursesTeachers                 DpoCoursesTeachers        `bun:"rel:has-many" json:"dpoCoursesTeachers"`
	DpoCoursesTeachersForDelete        []uuid.UUID               `bun:"-" json:"dpoCoursesTeachersForDelete"`
	DpoCoursesDates                    DpoCoursesDates           `bun:"rel:has-many" json:"dpoCoursesDates"`
	DpoCoursesDatesForDelete           []uuid.UUID               `bun:"-" json:"dpoCoursesDatesForDelete"`
	FormPattern                        *FormPattern              `bun:"rel:belongs-to" json:"formPattern"`
	FormPatternID                      uuid.NullUUID             `bun:"type:uuid" json:"formPatternId"`
}

type DpoCourses []*DpoCourse

func (item *DpoCourse) SetIdForChildren() {
	for i := range item.DpoCoursesTeachers {
		item.DpoCoursesTeachers[i].DpoCourseID = item.ID
	}
	for i := range item.DpoCoursesSpecializations {
		item.DpoCoursesSpecializations[i].DpoCourseID = item.ID
	}
	for i := range item.DpoCoursesDates {
		item.DpoCoursesDates[i].DpoCourseID = item.ID
	}
}

func (item *DpoCourse) SetForeignKeys() {
	item.FormPatternID = item.FormPattern.ID
}

func (item *DpoCourse) SetFilePath(fileID *string) *string {
	return item.FormPattern.SetFilePath(fileID)
}
