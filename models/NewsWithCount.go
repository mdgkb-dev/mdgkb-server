package models

type NewsWithCount struct {
	News  []*News `json:"items"`
	Count int     `json:"count"`
}
