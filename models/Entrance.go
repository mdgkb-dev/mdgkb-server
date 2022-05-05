package models

import "github.com/google/uuid"

type Entrance struct {
	ID         uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Building   *Building     `bun:"rel:belongs-to" json:"building"`
	BuildingId uuid.UUID     `bun:"type:uuid" json:"buildingId"`
	Number     int           `bun:"type:integer" json:"number"`
	Divisions  []*Division   `bun:"rel:has-many" json:"divisions"`
}

type Entrances []*Entrance
