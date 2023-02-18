package models

type DivisionsWithCount struct {
	Divisions Divisions `json:"items"`
	Count     int       `json:"count"`
}
