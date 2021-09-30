package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Menu struct {
	bun.BaseModel `bun:"menus,alias:menus"`
	ID                  uuid.UUID      `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name               string         `json:"name"`
	Link                string         `json:"link"`
	Top bool `json:"top"`
	Side bool `json:"side"`
	Page   *Page `bun:"rel:belongs-to" json:"page"`
	PageId uuid.UUID    `bun:"type:uuid" json:"PageId"`

	SubMenus SubMenus `bun:"rel:has-many" json:"subMenus"`
	SubMenusForDelete []string `bun:"-" json:"subMenusForDelete"`
}

type Menus []*Menu

func (item *Menu) SetIdForChildren() {
	if len(item.SubMenus) == 0 {
		return
	}
	for i := range item.SubMenus {
		item.SubMenus[i].MenuId = item.ID
	}
}