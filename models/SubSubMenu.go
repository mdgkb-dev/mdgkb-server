package models

import (
	"mdgkb/mdgkb-server/helpers/uploadHelper"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type SubSubMenu struct {
	bun.BaseModel `bun:"sub_sub_menus,alias:sub_sub_menus"`
	ID            uuid.UUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Link          string        `json:"link"`
	SubMenu       *SubMenu      `bun:"rel:belongs-to" json:"subMenu"`
	SubMenuId     uuid.UUID     `bun:"type:uuid" json:"subMenuId"`
	Icon          *FileInfo     `bun:"rel:belongs-to" json:"icon"`
	IconId        uuid.NullUUID `bun:"type:uuid"  json:"iconId"`
	Order         uint          `bun:"sub_sub_menu_order" json:"order"`

	Page   *Page         `bun:"rel:belongs-to" json:"page"`
	PageId uuid.NullUUID `bun:"type:uuid" json:"PageId"`
}

type SubSubMenus []*SubSubMenu

func (items SubSubMenus) SetFilePath(fileID *string) *string {
	for _, item := range items {
		if item.Icon.ID.UUID.String() == *fileID {
			item.Icon.FileSystemPath = uploadHelper.BuildPath(fileID)
			return &item.Icon.FileSystemPath
		}
	}
	return nil
}

func (items SubSubMenus) SetForeignKeys() {
	for i := range items {
		items[i].IconId = items[i].Icon.ID
	}
}

func (items SubSubMenus) GetFileInfos() FileInfos {
	itemsForGet := make(FileInfos, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.Icon)
	}
	return itemsForGet
}
