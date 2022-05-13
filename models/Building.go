package models

import "github.com/google/uuid"

type Building struct {
	ID        uuid.UUID   `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name      string      `json:"name"`
	Address   string      `json:"address"`
	Number    string      `json:"number"`
	Floors    []*Floor    `bun:"rel:has-many" json:"floors"`
	Entrances []*Entrance `bun:"rel:has-many" json:"entrances"`
}
