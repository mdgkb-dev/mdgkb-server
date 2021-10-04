package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/helpers/uploadHelper"
)

type SubMenu struct {
	bun.BaseModel `bun:"sub_menus,alias:sub_menus"`
	ID                  uuid.UUID      `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name               string         `json:"name"`
	Link                string         `json:"link"`
	Menu   *Menu `bun:"rel:belongs-to" json:"menu"`
	MenuId uuid.UUID    `bun:"type:uuid" json:"menuId"`

	Page   *Page `bun:"rel:belongs-to" json:"page"`
	PageId uuid.NullUUID    `bun:"type:uuid" json:"PageId"`
	Icon *FileInfo `bun:"rel:belongs-to" json:"icon"`
	IconId uuid.NullUUID `bun:"type:uuid"  json:"iconId"`

	SubSubMenus SubSubMenus `bun:"rel:has-many" json:"subSubMenus"`
	SubSubMenusForDelete []string `bun:"-" json:"subSubMenusForDelete"`
}

type SubMenus []*SubMenu

func (item *SubMenu) SetIdForChildren() {
	if len(item.SubSubMenus) == 0 {
		return
	}
	for i := range item.SubSubMenus {
		item.SubSubMenus[i].SubMenuId = item.ID
	}
}

func (items SubMenus)SetIdForChildren() {
	for i := range items {
		items[i].SetIdForChildren()
	}
}


func (items SubMenus) GetIDForDelete() []string {
	idPool := make([]string, 0)
	for _, item := range items {
		idPool = append(idPool, item.SubSubMenusForDelete... )
	}
	return idPool
}

func (items SubMenus) GetSubSubMenus() SubSubMenus {
	itemsForGet := make(SubSubMenus, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.SubSubMenus...)
	}
	return itemsForGet
}

func (items SubMenus) SetFilePath(fileId *string) *string {
	for _, item := range items {
		if item.Icon.ID.UUID.String() == *fileId {
			item.Icon.FileSystemPath = uploadHelper.BuildPath(fileId)
			return &item.Icon.FileSystemPath
		}
		path := item.SubSubMenus.SetFilePath(fileId)
		if path != nil {
			return path
		}
	}
	return nil
}

func (items SubMenus) SetForeignKeys()  {
	for i := range items {
		items[i].IconId.UUID = items[i].Icon.ID.UUID
		items[i].IconId = items[i].Icon.ID
	}
}


func (items SubMenus) GetFileInfos() FileInfos {
	itemsForGet := make(FileInfos, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.Icon)
	}
	return itemsForGet
}