package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Page struct {
	bun.BaseModel `bun:"pages,alias:pages"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Title         string        `json:"title"`
	Content       string        `json:"content"`
	Slug          string        `json:"slug"`
	Link          string        `json:"link"`
	WithComments  bool          `json:"withComments"`

	PagesGroup string `json:"pagesGroup"`

	PageSideMenus          PageSideMenus `bun:"rel:has-many" json:"pageSideMenus"`
	PageSideMenusForDelete []uuid.UUID   `bun:"-" json:"pageSideMenusForDelete"`

	PageSections          PageSections `bun:"rel:has-many" json:"pageSections"`
	PageSectionsForDelete []string     `bun:"-" json:"pageSectionsForDelete"`

	PageComments          PageComments `bun:"rel:has-many" json:"pageComments"`
	PageCommentsForDelete []string     `bun:"-" json:"pageCommentsForDelete"`

	PageDocuments          PageDocuments `bun:"rel:has-many" json:"pageDocuments"`
	PageDocumentsForDelete []string      `bun:"-" json:"pageDocumentsForDelete"`

	PageImages          PageImages  `bun:"rel:has-many" json:"pageImages"`
	PageImagesForDelete []uuid.UUID `bun:"-" json:"pageImagesForDelete"`
}

type Pages []*Page

type PagesWithCount struct {
	Pages Pages `json:"items"`
	Count int   `json:"count"`
}

func (item *Page) SetIDForChildren() {
	if len(item.PageComments) > 0 {
		for i := range item.PageComments {
			item.PageComments[i].PageID = item.ID
		}
	}
	for i := range item.PageDocuments {
		item.PageDocuments[i].PageID = item.ID
	}
	for i := range item.PageSideMenus {
		item.PageSideMenus[i].PageID = item.ID
	}
	for i := range item.PageImages {
		item.PageImages[i].PageID = item.ID
	}
}

func (item *Page) SetFilePath(fileID *string) *string {
	path := item.PageSideMenus.SetFilePath(fileID)
	if path != nil {
		return path
	}
	path = item.PageDocuments.SetFilePath(fileID)
	if path != nil {
		return path
	}
	path = item.PageImages.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}
