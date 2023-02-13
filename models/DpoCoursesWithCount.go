package models

type DpoCoursesWithCount struct {
	DpoCourses DpoCourses `json:"items"`
	Count      int        `json:"count"`
}
