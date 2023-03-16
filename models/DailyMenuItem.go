package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DailyMenuItem struct {
	bun.BaseModel     `bun:"daily_menu_items,alias:daily_menu_items"`
	ID                uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name              string        `json:"name"`
	Price             uint          `json:"price"`
	Caloric           uint          `json:"caloric"`
	Weight            uint          `json:"weight"`
	AdditionalWeight  uint          `json:"additionalWeight"`
	Order             uint          `bun:"item_order" json:"order"`
	Quantity          int           `json:"quantity"`
	DailyMenu         *DailyMenu    `bun:"rel:belongs-to" json:"dailyMenu"`
	DailyMenuID       uuid.NullUUID `bun:"type:uuid"  json:"dailyMenuId"`
	DishesGroup       *DishesGroup  `bun:"rel:belongs-to" json:"dishesGroup"`
	DishesGroupID     uuid.NullUUID `bun:"type:uuid"  json:"dishesGroupId"`
	DishSample        *DishSample   `bun:"rel:belongs-to" json:"dishSample"`
	DishSampleID      uuid.NullUUID `bun:"type:uuid"  json:"dishSampleId"`
	Available         bool          `json:"available"`
	FromOtherMenu     bool          `json:"fromOtherMenu"`
	Cook              bool          `json:"cook"`
	TomorrowAvailable bool          `json:"tomorrowAvailable"`

	Proteins      uint `json:"proteins"`
	Fats          uint `json:"fats"`
	Carbohydrates uint `json:"carbohydrates"`

	Dietary bool `json:"dietary"`
	Lean    bool `json:"lean"`
}

type DailyMenuItems []*DailyMenuItem
