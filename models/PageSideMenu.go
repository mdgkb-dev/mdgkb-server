package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PageSideMenu struct {
	bun.BaseModel         `bun:"page_side_menus,alias:page_side_menus"`
	ID                    uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name                  string        `json:"name"`
	RouteAnchor           string        `json:"routeAnchor"`
	Description           string        `json:"description"`
	Order                 int           `bun:"item_order" json:"order"`
	PageSections          PageSections  `bun:"rel:has-many" json:"pageSections"`
	PageSectionsForDelete []uuid.UUID   `bun:"-" json:"pageSectionsForDelete"`

	Page   *Page         `bun:"rel:belongs-to" json:"page"`
	PageID uuid.NullUUID `bun:"type:uuid" json:"pageId"`
}

type PageSideMenus []*PageSideMenu

func (item PageSideMenu) SetFilePath(fileID *string) *string {
	for i := range item.PageSections {
		filePath := item.PageSections[i].SetFilePath(fileID)
		if filePath != nil {
			return filePath
		}
	}
	return nil
}

func (item *PageSideMenu) SetIDForChildren() {
	for i := range item.PageSections {
		item.PageSections[i].PageSideMenuID = item.ID
	}
}

func (items PageSideMenus) SetIDForChildren() {
	for i := range items {
		items[i].SetIDForChildren()
	}
}

func (items PageSideMenus) GetPageSections() PageSections {
	itemsForGet := make(PageSections, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PageSections...)
	}
	return itemsForGet
}

func (items PageSideMenus) GetPageSectionsForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PageSectionsForDelete...)
	}
	return itemsForGet
}
