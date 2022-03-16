package models

type SearchModel struct {
	Query         string       `json:"query"`
	SearchGroupID string       `json:"searchGroupId"`
	SearchGroups  SearchGroups `json:"searchGroups"`
	SearchGroup   *SearchGroup `json:"searchGroup"`
}

func (item *SearchModel) ParseMap(re map[string]interface{}) {
	item.SearchGroup.ParseMap(re)
}
