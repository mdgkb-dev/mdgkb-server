package models

type UsersWithCount struct {
	Users Users `json:"users"`
	Count int   `json:"count"`
}
