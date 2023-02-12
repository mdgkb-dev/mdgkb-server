package models

type ResidencyCoursesWithCount struct {
	ResidencyCourses ResidencyCourses `json:"items"`
	Count            int              `json:"count"`
}
