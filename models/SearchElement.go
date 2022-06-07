package models

import "github.com/google/uuid"

type SearchElement struct {
	ID                 uuid.UUID          `json:"id"`
	Description        string             `json:"description"`
	Value              string             `json:"value"`
	Label              string             `json:"label"`
	Route              string             `json:"route"`
	SearchGroup        *SearchGroup       `json:"searchGroup"`
	SearchElementMetas SearchElementMetas `json:"searchElementMetas"`
}

type SearchElements []*SearchElement
