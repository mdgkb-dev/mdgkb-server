package models

import (
	"github.com/google/uuid"
)

type ChartQuery struct {
	EntityID uuid.UUID     `bun:"type=uuid" json:"entityId"`
	Type     ChartDataType `json:"chartDataType"`
	DataSets ChartDataSets `json:"dataSets"`
}

type ChartDataType string

const (
	ChartDataTypeNewsDatesViews  ChartDataType = "newsDatesViews"
	ChartDataTypeNewsCitiesViews ChartDataType = "newsCitiesViews"
)

func (item *ChartQuery) ParseExportOptions(map[string]map[string]interface{}) error {
	return nil
}
