package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Diet struct {
	bun.BaseModel     `bun:"diets,alias:diets"`
	ID                uuid.NullUUID `bun:",pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name              string        `json:"name"`
	ShortName         string        `json:"shortName"`
	SiteName          string        `json:"siteName"`
	Diabetes          bool          `json:"diabetes"`
	Order             uint          `bun:"diet_order" json:"order"`
	MotherDiet        *MotherDiet   `bun:"rel:belongs-to" json:"motherDiet"`
	MotherDietID      uuid.NullUUID `bun:"type:uuid"  json:"motherDietId"`
	DietGroup         *DietGroup    `bun:"rel:belongs-to" json:"dietGroup"`
	DietGroupID       uuid.NullUUID `bun:"type:uuid"  json:"dietGroupId"`
	DietAges          DietAges      `bun:"rel:has-many,join:id=diet_id" json:"dietAges"`
	DietAgesForDelete []uuid.UUID   `bun:"-" json:"dietAgesForDelete"`
}

type MotherDiet Diet
type Diets []Diet
