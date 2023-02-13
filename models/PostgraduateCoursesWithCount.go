package models

type PostgraduateCoursesWithCount struct {
	PostgraduateCourses PostgraduateCourses `json:"items"`
	Count               int                 `json:"count"`
}
