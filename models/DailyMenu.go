package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DailyMenu struct {
	bun.BaseModel  `bun:"daily_menus,alias:daily_menus"`
	ID             uuid.NullUUID  `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Date           time.Time      `bun:"daily_menu_date" json:"date"`
	DailyMenuItems DailyMenuItems `bun:"rel:has-many" json:"dailyMenuItems"`
}

type DailyMenus []*DailyMenu
