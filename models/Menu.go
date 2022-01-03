package models

import (
	"mdgkb/mdgkb-server/helpers/uploadHelper"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Menu struct {
	bun.BaseModel `bun:"menus,alias:menus"`
	ID            uuid.UUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Link          string        `json:"link"`
	Top           bool          `json:"top"`
	Side          bool          `json:"side"`
	Order         uint          `bun:"menu_order" json:"order"`
	Page          *Page         `bun:"rel:belongs-to" json:"page"`
	PageId        uuid.NullUUID `bun:"type:uuid" json:"pageId"`
	Icon          *FileInfo     `bun:"rel:belongs-to" json:"icon"`
	IconId        uuid.NullUUID `bun:"type:uuid"  json:"iconId"`

	SubMenus          SubMenus    `bun:"rel:has-many" json:"subMenus"`
	SubMenusForDelete []uuid.UUID `bun:"-" json:"subMenusForDelete"`
}

type Menus []*Menu

func (items Menus) GetIcons() FileInfos {
	itemsForGet := make(FileInfos, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.Icon)
	}
	return itemsForGet
}

func (items Menus) GetSubMenus() SubMenus {
	itemsForGet := make(SubMenus, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.SubMenus...)
	}
	return itemsForGet
}

func (items Menus) GetSubMenusForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.SubMenusForDelete...)
	}
	return itemsForGet
}

func (item *Menu) SetIdForChildren() {
	if len(item.SubMenus) == 0 {
		return
	}
	for i := range item.SubMenus {
		item.SubMenus[i].MenuId = item.ID
	}
}

func (items Menus) SetIdForChildren() {
	for i := range items {
		items[i].SetIdForChildren()
	}
}

func (item *Menu) SetFilePath(fileID *string) *string {
	if item.Icon.ID.UUID.String() == *fileID {
		item.Icon.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.Icon.FileSystemPath
	}
	path := item.SubMenus.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}

func (item *Menu) SetForeignKeys() {
	item.IconId = item.Icon.ID
}

func (items Menus) SetForeignKeys() {
	for i := range items {
		items[i].SetForeignKeys()
	}
}
