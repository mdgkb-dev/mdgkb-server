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
	Number        *uint         `bun:",autoincrement,notnull," json:"number"`

	FormValue   *FormValue    `bun:"rel:belongs-to" json:"formValue"`
	FormValueID uuid.NullUUID `bun:"type:uuid" json:"formValueId"`

	DailyMenuOrderItems          DailyMenuOrderItems `bun:"rel:has-many" json:"dailyMenuOrderItems"`
	DailyMenuOrderItemsForDelete []uuid.UUID         `bun:"-" json:"dailyMenuOrderItemsForDelete"`

	IsNew        string        `bun:"-" json:"isNew"`
	CreatedAt    string        `bun:"-" json:"createdAt"`
	FormStatusID string        `bun:"-" json:"formStatusId"`
	User         *User         `bun:"rel:belongs-to" json:"user"`
	UserID       uuid.NullUUID `bun:"user_id,nullzero,type:uuid" json:"userId"`
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
