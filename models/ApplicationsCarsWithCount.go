package models

type ApplicationsCarsWithCount struct {
	ApplicationsCars ApplicationsCars `json:"applicationsCars"`
	Count            int              `json:"count"`
}
