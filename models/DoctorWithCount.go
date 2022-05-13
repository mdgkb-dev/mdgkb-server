package models

type DoctorsWithCount struct {
	Doctors Doctors `json:"doctors"`
	Count   int     `json:"count"`
}
