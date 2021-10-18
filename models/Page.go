package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Page struct {
	bun.BaseModel `bun:"pages,alias:pages"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	Slug          string    `json:"slug"`
	Link          string    `json:"link"`
	WithComments  bool      `json:"withComments"`

	PageComments          PageComments `bun:"rel:has-many" json:"pageComments"`
	PageCommentsForDelete []string     `bun:"-" json:"pageCommentsForDelete"`

	PageDocuments          PageDocuments `bun:"rel:has-many" json:"pageDocuments"`
	PageDocumentsForDelete []string      `bun:"-" json:"pageDocumentsForDelete"`

	PageImages          PageImages  `bun:"rel:has-many" json:"pageImages"`
	PageImagesForDelete []uuid.UUID `bun:"-" json:"pageImagesForDelete"`
}

type Pages []*Page

func (item *Page) SetIdForChildren() {
	if len(item.PageComments) < 0 {
		return
	}
	for i := range item.PageComments {
		item.PageComments[i].PageId = item.ID
	}
	for i := range item.PageDocuments {
		item.PageDocuments[i].PageID = item.ID
	}
	for i := range item.PageImages {
		item.PageImages[i].PageID = item.ID
	}
}

func (item *Page) SetFilePath(fileId *string) *string {
	path := item.PageDocuments.SetFilePath(fileId)
	if path != nil {
		return path
	}
	path = item.PageImages.SetFilePath(fileId)
	if path != nil {
		return path
	}
	return nil
}
