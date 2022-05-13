package models

import (
	"github.com/uptrace/bun"
)

type ApplicationsCount struct {
	bun.BaseModel `bun:"applications_counts,alias:applications_counts"`
	TableName     string `json:"tableName"`
	Count         int    `json:"count"`
}

type ApplicationsCounts []*ApplicationsCount
