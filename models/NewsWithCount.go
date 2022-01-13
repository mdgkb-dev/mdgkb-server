package models

type NewsWithCount struct {
	News  []*News `json:"news"`
	Count int     `json:"count"`
}
