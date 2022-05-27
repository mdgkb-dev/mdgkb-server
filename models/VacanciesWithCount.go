package models

type VacanciesWithCount struct {
	Vacancies Vacancies `json:"vacancies"`
	Count     int       `json:"count"`
}
