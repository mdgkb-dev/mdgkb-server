package models

type PostgraduateCoursesWithCount struct {
	PostgraduateCourses PostgraduateCourses `json:"postgraduateCourses"`
	Count               int                 `json:"count"`
}
