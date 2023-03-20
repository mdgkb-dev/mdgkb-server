package models

type VisitsApplicationsWithCount struct {
	VisitsApplications VisitsApplications `json:"items"`
	Count              int                `json:"count"`
}
