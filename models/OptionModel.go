package models

type OptionModel struct {
	TableName string `json:"tableName"`
	Value     string `json:"value"`
	Label     string `json:"label"`
}

type OptionModels []*OptionModel
