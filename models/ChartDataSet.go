package models

type ChartDataSet struct {
	Value float64 `json:"value"`
	Label string  `json:"label"`
}

type ChartDataSets []*ChartDataSet
