package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type TreatDirection struct {
	bun.BaseModel `bun:"treat_directions,alias:treat_directions"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Divisions     Divisions     `bun:"rel:has-many" json:"divisions"`
}

type TreatDirections []*TreatDirection
