package models

type SupportMessagesWithCount struct {
	SupportMessages SupportMessages `json:"supportMessages"`
	Count           int             `json:"count"`
}
