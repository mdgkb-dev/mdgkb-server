package models

type DivisionsWithCount struct {
	Divisions Divisions `json:"divisions"`
	Count     int       `json:"count"`
}
