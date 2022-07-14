package models

type VisitsApplicationsWithCount struct {
	VisitsApplications VisitsApplications `json:"visitsApplications"`
	Count              int                `json:"count"`
}
