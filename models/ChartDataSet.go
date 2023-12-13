package models

type ChartDataSet struct {
	Value float64 `json:"value"`
	Label string  `json:"label"`

	Data []float64 `json:"data"`
}

type ChartDataSets []*ChartDataSet
