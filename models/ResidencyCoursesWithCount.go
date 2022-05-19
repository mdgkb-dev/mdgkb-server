package models

type ResidencyCoursesWithCount struct {
	ResidencyCourses ResidencyCourses `json:"residencyCourses"`
	Count            int              `json:"count"`
}
