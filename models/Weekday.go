package models

import "github.com/google/uuid"

type Weekday struct {
	ID        uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Number    int       `json:"number"`
	Name      string    `json:"name"`
	ShortName string    `json:"shortName"`
}

type Weekdays []*Weekday
