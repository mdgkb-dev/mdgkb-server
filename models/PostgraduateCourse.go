package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PostgraduateCourse struct {
	bun.BaseModel                               `bun:"postgraduate_courses,alias:postgraduate_courses"`
	ID                                          uuid.NullUUID                      `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	Description                                 string                             `json:"description"`
	PostgraduateCoursesSpecializations          PostgraduateCoursesSpecializations `bun:"rel:has-many" json:"postgraduateCoursesSpecializations"`
	PostgraduateCoursesSpecializationsForDelete []uuid.UUID                        `bun:"-" json:"postgraduateCoursesSpecializationsForDelete"`
	PostgraduateCoursesTeachers                 PostgraduateCoursesTeachers        `bun:"rel:has-many" json:"postgraduateCoursesTeachers"`
	PostgraduateCoursesTeachersForDelete        []uuid.UUID                        `bun:"-" json:"postgraduateCoursesForDelete"`
	PostgraduateCoursesDates                    PostgraduateCoursesDates           `bun:"rel:has-many" json:"postgraduateCoursesDates"`
	PostgraduateCoursesDatesForDelete           []uuid.UUID                        `bun:"-" json:"postgraduateCoursesDatesForDelete"`
}

type PostgraduateCourses []*PostgraduateCourse

func (item *PostgraduateCourse) SetIdForChildren() {
	for i := range item.PostgraduateCoursesTeachers {
		item.PostgraduateCoursesTeachers[i].PostgraduateCourseID = item.ID
	}
	for i := range item.PostgraduateCoursesSpecializations {
		item.PostgraduateCoursesSpecializations[i].PostgraduateCourseID = item.ID
	}
	for i := range item.PostgraduateCoursesDates {
		item.PostgraduateCoursesDates[i].PostgraduateCourseID = item.ID
	}
}
