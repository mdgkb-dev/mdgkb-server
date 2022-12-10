package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DailyMenu struct {
	bun.BaseModel  `bun:"daily_menus,alias:daily_menus"`
	ID             uuid.NullUUID  `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Date           time.Time      `bun:"item_date" json:"date"`
	Name           string         `json:"name"`
	Order          uint8          `bun:"item_order" json:"order"`
	DailyMenuItems DailyMenuItems `bun:"rel:has-many" json:"dailyMenuItems"`
}

type DailyMenus []*DailyMenu

func (item *DailyMenu) SetIDForChildren() {
	for i := range item.DailyMenuItems {
		item.DailyMenuItems[i].DailyMenuID = item.ID
	}
}
