package models

type QuestionsWithCount struct {
	Questions Questions `json:"items"`
	Count     int       `json:"count"`
}
