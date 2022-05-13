package models

import "github.com/google/uuid"

type Category struct {
	ID   uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name string    `json:"name"`
}

type Categories []*Category
