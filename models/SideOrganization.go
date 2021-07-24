package models

import "github.com/google/uuid"

type SideOrganization struct {
	ID      uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name    string    `json:"name"`
	Site    string    `json:"site"`
	Phone   string    `json:"phone"`
	Address string    `json:"address"`
}
