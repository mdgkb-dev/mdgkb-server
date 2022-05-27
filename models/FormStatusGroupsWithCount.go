package models

type FormStatusGroupsWithCount struct {
	FormStatusGroups FormStatusGroups `json:"formStatusGroups"`
	Count            int              `json:"count"`
}
