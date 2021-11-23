package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type SearchGroup struct {
	bun.BaseModel    `bun:"search_groups,alias:search_groups"`
	ID               uuid.UUID        `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	Label string `json:"label"`
	Order int `bun:"search_group_order" json:"order"`
	Route string `json:"route"`
	Table string `bun:"search_group_table" json:"table"`
	SearchColumn string `json:"searchColumn"`
	LabelColumn string `json:"labelColumn"`
	ValueColumn string `json:"valueColumnColumn"`

	SearchElements SearchElements `bun:"-" json:"options"`
}

type SearchGroups []*SearchGroup

func (item *SearchGroup) BuildRoutes()  {
	for i := range item.SearchElements {
		item.SearchElements[i].Value = fmt.Sprintf("%s/%s", item.Route, item.SearchElements[i].Value)
	}
}