package models

import (
	"github.com/AvraamMavridis/randomcolor"
	"github.com/google/uuid"
)

type ChartQuery struct {
	EntityID uuid.NullUUID `bun:"type=uuid" json:"entityId"`
	Type     ChartDataType `json:"chartDataType"`
	DataSets ChartDataSets `json:"datasets"`
	Labels   []string      `json:"labels"`
}

type ChartDataType string

const (
	ChartDataTypeNewsDatesViews  ChartDataType = "newsDatesViews"
	ChartDataTypeNewsCitiesViews ChartDataType = "newsCitiesViews"
)

func (item *ChartQuery) ParseExportOptions(map[string]map[string]interface{}) error {
	return nil
}

func (item *ChartQuery) InitFromDataSets(dataSets ChartDataSets) {
	item.DataSets = append(item.DataSets, &ChartDataSet{})
	for _, dataSet := range dataSets {
		item.Labels = append(item.Labels, dataSet.Label)
		item.DataSets[0].Data = append(item.DataSets[0].Data, dataSet.Value)
		item.DataSets[0].BackgroundColor = append(item.DataSets[0].BackgroundColor, randomcolor.GetRandomColorInHex())
	}
}
