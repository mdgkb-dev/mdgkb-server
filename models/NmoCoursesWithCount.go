package models

type NmoCoursesWithCount struct {
	NmoCourses NmoCourses `json:"items"`
	Count      int        `json:"count"`
}
