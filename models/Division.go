package models

import "github.com/google/uuid"

type Division struct {
	ID      uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name    string    `json:"name"`
	Info    string    `json:"info"`
	Phone   string    `json:"phone"`
	Email   string    `json:"email"`
	Address string    `json:"address"`
	Slug    string    `json:"slug"`
	Doctors  []*Doctor `bun:"rel:has-many" json:"doctors"`
	FloorId uuid.UUID `bun:"type:uuid" json:"floorId"`
}
