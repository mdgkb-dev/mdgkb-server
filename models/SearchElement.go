package models

type SearchElement struct {
	Description string `json:"description"`
	Value       string `json:"value"`
	Label       string `json:"label"`
}

type SearchElements []*SearchElement
