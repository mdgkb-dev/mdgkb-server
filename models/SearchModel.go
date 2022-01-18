package models

type SearchModel struct {
	Query         string       `json:"query"`
	SearchGroupID string       `json:"searchGroupId"`
	SearchGroups  SearchGroups `json:"searchGroups"`
	SearchGroup   *SearchGroup `json:"searchGroup"`
	Mode          SearchMode   `json:"searchMode"`
}

type SearchMode string

const (
	SearchModeMain    SearchMode = "main"
	SearchModeObjects            = "objects"
)
