package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DailyMenuOrder struct {
	bun.BaseModel `bun:"daily_menu_orders,alias:daily_menu_orders"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Date          time.Time     `bun:"item_date" json:"date"`
	BoxNumber     uint8         `json:"boxNumber"`
	Number        uint          `json:"number"`
	Price         uint          `json:"price"`

	FormValue   *FormValue    `bun:"rel:belongs-to" json:"formValue"`
	FormValueID uuid.NullUUID `bun:"type:uuid" json:"formValueId"`

	DailyMenuOrderItems DailyMenuOrderItems `bun:"rel:has-many" json:"dailyMenuOrderItems"`
}

type DailyMenuOrders []*DailyMenuOrder
