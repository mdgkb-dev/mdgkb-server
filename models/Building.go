package models

import "github.com/google/uuid"

type Building struct {
	ID      uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name    string    `json:"name"`
	Address string    `json:"address"`
	//Status      string    `json:"status"`
	Description string   `json:"description"`
	Floors      []*Floor `bun:"rel:has-many" json:"floors"`
}
