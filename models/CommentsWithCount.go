package models

type CommentsWithCount struct {
	Comments Comments `json:"comments"`
	Count    int      `json:"count"`
}
