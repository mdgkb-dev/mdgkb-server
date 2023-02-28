package models

import (
	"fmt"

	"github.com/pro-assistance/pro-assister/sqlHelper/paginator"
)

type SearchModel struct {
	Suggester       bool                `json:"suggester"`
	SearchElements  SearchElements      `bun:"-" json:"options"`
	Query           string              `json:"query"`
	MustBeTranslate bool                `json:"mustBeTranslate"`
	TranslitQuery   string              `json:"translitQuery"`
	SearchGroupID   string              `json:"searchGroupId"`
	SearchGroups    SearchGroups        `json:"searchGroups"`
	SearchGroup     *SearchGroup        `json:"searchGroup"`
	SearchColumn    string              `json:"searchColumn"`
	Count           int                 `json:"count"`
	Pagination      paginator.Paginator `json:"pagination"`
}

func (item *SearchModel) BuildRoutes() {
	for i := range item.SearchGroups {
		item.SearchGroups[i].BuildRoutes()
	}
	for i := range item.SearchElements {
		group := item.findGroupByKey(item.SearchElements[i].Key)
		item.SearchElements[i].Route = fmt.Sprintf("%s/%s", group.Route, item.SearchElements[i].Value)
	}
}

func (item *SearchModel) findGroupByKey(key string) SearchGroup {
	group := SearchGroup{}
	for _, searchGroup := range item.SearchGroups {
		if searchGroup.Key == key {
			group = *searchGroup
			break
		}
	}
	return group
}
