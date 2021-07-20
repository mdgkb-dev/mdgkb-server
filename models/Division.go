package models

import "github.com/google/uuid"

type Division struct {
	ID      uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name    string    `json:"name"`
	Status  string    `json:"status"`
	Phone   string    `json:"phone"`
	Email   string    `json:"email"`
	Address string    `json:"address"`
	Slug    string    `json:"slug"`
	FloorId uuid.UUID `bun:"type:uuid"`
}
