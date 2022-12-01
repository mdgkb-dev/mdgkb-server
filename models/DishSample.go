package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DishSample struct {
	bun.BaseModel `bun:"dishes_samples,alias:dish_samples"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Price         uint          `json:"price"`
	Caloric       uint          `json:"caloric"`
	Weight        uint          `json:"weight"`
	DishesGroup   *DishesGroup  `bun:"rel:belongs-to" json:"dishesGroup"`
	DishesGroupID uuid.NullUUID `bun:"type:uuid"  json:"dishesGroupId"`
}

type DishSamples []*DishSample
