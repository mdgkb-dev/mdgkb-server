package models

type OptionModel struct {
	TableName  string `json:"tableName"`
	SortColumn string `json:"sortColumn"`
	Value      string `json:"value"`
	Label      string `json:"label"`
}

type OptionModels []*OptionModel
