package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type SubSubMenu struct {
	bun.BaseModel `bun:"sub_sub_menus,alias:sub_sub_menus"`
	ID                  uuid.UUID      `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name               string         `json:"name"`
	Link                string         `json:"link"`
	SubMenu   *SubMenu `bun:"rel:belongs-to" json:"subMenu"`
	SubMenuId uuid.UUID    `bun:"type:uuid" json:"subMenuId"`

	Page   *Page `bun:"rel:belongs-to" json:"page"`
	PageId uuid.UUID    `bun:"type:uuid" json:"PageId"`
}

type SubSubMenus []*SubSubMenu