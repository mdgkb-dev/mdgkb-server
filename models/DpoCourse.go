package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DpoCourse struct {
	bun.BaseModel                      `bun:"dpo_courses,alias:dpo_courses"`
	ID                                 uuid.NullUUID             `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	Name                               string                    `json:"name"`
	Description                        string                    `json:"description"`
	Order                              int                       `bun:"dpo_course_order" json:"order"`
	IsNmo                              bool                      `json:"isNmo"`
	LinkNmo                            string                    `json:"linkNmo"`
	Listeners                          int                       `json:"listeners"`
	Hours                              int                       `json:"hours"`
	Cost                               int                       `json:"cost"`
	DpoCoursesSpecializations          DpoCoursesSpecializations `bun:"rel:has-many" json:"dpoCoursesSpecializations"`
	DpoCoursesSpecializationsForDelete []uuid.UUID               `bun:"-" json:"dpoCoursesSpecializationsForDelete"`
	DpoCoursesTeachers                 DpoCoursesTeachers        `bun:"rel:has-many" json:"dpoCoursesTeachers"`
	DpoCoursesTeachersForDelete        []uuid.UUID               `bun:"-" json:"dpoCoursesForDelete"`
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
