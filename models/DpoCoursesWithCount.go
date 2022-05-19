package models

type DpoCoursesWithCount struct {
	DpoCourses DpoCourses `json:"dpoCourses"`
	Count      int        `json:"count"`
}
