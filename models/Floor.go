package models

import "github.com/google/uuid"

type Floor struct {
	ID         uuid.UUID   `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	BuildingId uuid.UUID   `bun:"type:uuid"`
	Number     int8        `bun:"type:integer"`
	Divisions  []*Division `bun:"rel:has-many" json:"divisions"`
}
