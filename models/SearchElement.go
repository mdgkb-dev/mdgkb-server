package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type SearchElement struct {
	bun.BaseModel      `bun:"search_elements,alias:search_elements"`
	ID                 uuid.UUID          `json:"id"`
	Description        string             `json:"description"`
	Value              string             `json:"value"`
	Label              string             `json:"label"`
	Route              string             `json:"route"`
	Key                string             `json:"key"`
	SearchGroup        *SearchGroup       `json:"searchGroup"`
	SearchElementMetas SearchElementMetas `json:"searchElementMetas"`
	Rank               float32            `json:"-"`
}

type SearchElements []*SearchElement
