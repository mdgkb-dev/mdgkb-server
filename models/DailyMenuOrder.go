package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DailyMenuOrder struct {
	bun.BaseModel `bun:"daily_menu_orders,select:daily_menu_orders_view,alias:daily_menu_orders_view"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Date          time.Time     `bun:"item_date" json:"date"`
	BoxNumber     uint          `json:"boxNumber"`
	Number        uint          `bun:",notnull,autoincrement" json:"number"`

	FormValue   *FormValue    `bun:"rel:belongs-to" json:"formValue"`
	FormValueID uuid.NullUUID `bun:"type:uuid" json:"formValueId"`

	DailyMenuOrderItems          DailyMenuOrderItems `bun:"rel:has-many" json:"dailyMenuOrderItems"`
	DailyMenuOrderItemsForDelete []uuid.UUID         `bun:"-" json:"dailyMenuOrderItemsForDelete"`

	CreatedAt string `bun:"-" json:"createdAt"`
}

type DailyMenuOrders []*DailyMenuOrder

type DailyMenuOrdersWithCount struct {
	DailyMenuOrders DailyMenuOrders `json:"items"`
	Count           int             `json:"count"`
}

func (item *DailyMenuOrder) SetIDForChildren() {
	for i := range item.DailyMenuOrderItems {
		item.DailyMenuOrderItems[i].DailyMenuOrderID = item.ID
	}
}

func (item *DailyMenuOrder) SetForeignKeys() {
	item.FormValueID = item.FormValue.ID
}
