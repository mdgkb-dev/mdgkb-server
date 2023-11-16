package models

type VacanciesWithCount struct {
	Vacancies Vacancies `json:"items"`
	Count     int       `json:"count"`
}
