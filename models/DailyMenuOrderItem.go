package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DailyMenuOrderItem struct {
	bun.BaseModel `bun:"daily_menu_order_items,alias:daily_menu_order_items"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	Quantity uint8 `json:"quantity"`
	Price    uint8 `json:"price"`

	DailyMenuOrder   *DailyMenuOrder `bun:"rel:belongs-to" json:"dailyMenuOrder"`
	DailyMenuOrderID uuid.NullUUID   `bun:"type:uuid"  json:"dailyMenuOrderId"`

	DailyMenuItem   *DailyMenuItem `bun:"rel:belongs-to" json:"dailyMenuItem"`
	DailyMenuItemID uuid.NullUUID  `bun:"type:uuid"  json:"dailyMenuItemId"`
}

type DailyMenuOrderItems []*DailyMenuOrderItem
