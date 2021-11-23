package models

type SearchElement struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type SearchElements []*SearchElement