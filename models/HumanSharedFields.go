package models

type HumanSharedFields struct {
	FullName  string `bun:"-" json:"fullName"`
	DateBirth string `bun:"-" json:"dateBirth"`
	IsMale    string `bun:"-" json:"isMale"`
}
