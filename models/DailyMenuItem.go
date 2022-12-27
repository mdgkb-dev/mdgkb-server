package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DailyMenuItem struct {
	bun.BaseModel `bun:"daily_menu_items,alias:daily_menu_items"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Price         uint          `json:"price"`
	Caloric       uint          `json:"caloric"`
	Weight        uint          `json:"weight"`
	Order         uint          `bun:"item_order" json:"order"`
	DailyMenu     *DailyMenu    `bun:"rel:belongs-to" json:"dailyMenu"`
	DailyMenuID   uuid.NullUUID `bun:"type:uuid"  json:"dailyMenuId"`
	DishSample    *DishSample   `bun:"rel:belongs-to" json:"dishSample"`
	DishSampleID  uuid.NullUUID `bun:"type:uuid"  json:"dishSampleId"`
	Available     bool          `json:"available"`
	FromOtherMenu bool          `json:"fromOtherMenu"`
}

type DailyMenuItems []*DailyMenuItem
