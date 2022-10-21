package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DietGroup struct {
	bun.BaseModel `bun:"diets_groups,alias:diets_groups"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Diets         Diets         `bun:"rel:has-many" json:"diets"`
}

type DietsGroups []*DietGroup
