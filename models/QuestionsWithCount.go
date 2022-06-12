package models

type QuestionsWithCount struct {
	Questions Questions `json:"questions"`
	Count     int       `json:"count"`
}
