package httpHelper

import (
	"encoding/json"
	"fmt"
)

// SortModel model
type SortModel struct {
	Table *string `json:"table"`
	Col   *string `json:"col"`
	Order *Orders `json:"order"`
}

type SortModels []*SortModel

type Orders string

const (
	Asc  Orders = "asc"
	Desc Orders = "desc"
)

func (s *SortModel) GetTableAndCol() string {
	return fmt.Sprintf("%s.%s", *s.Table, *s.Col)
}

// ParseJSONToSortModel constructor
func ParseJSONToSortModel(args string) (sortModel SortModel, err error) {
	err = json.Unmarshal([]byte(args), &sortModel)
	if err != nil {
		return sortModel, err
	}
	return sortModel, err
}
