package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DishesGroup struct {
	bun.BaseModel `bun:"dishes_groups,alias:dishes_groups"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Order         uint          `bun:"dishes_group_order" json:"order"`
}

type DishesGroups []*DishesGroup
