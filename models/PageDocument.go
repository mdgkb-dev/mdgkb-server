package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/helpers/uploadHelper"
)

type PageDocument struct {
	bun.BaseModel `bun:"pages_documents,alias:pages_documents"`
	ID   uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`

	Page *Page `bun:"rel:belongs-to" json:"page"`
	PageId uuid.UUID `bun:"type:uuid" json:"pageId"`

	Document *Document `bun:"rel:belongs-to" json:"document"`
	DocumentID uuid.UUID `bun:"type:uuid" json:"documentId"`
}

type PageDocuments []*PageDocument

func (items PageDocuments) GetDocuments() Documents {
	itemsForGet := make(Documents, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.Document)
	}
	return itemsForGet
}

func (items PageDocuments) SetForeignKeys()  {
	for i := range items {
		items[i].DocumentID = items[i].Document.ID
	}
}

func (items PageDocuments) SetFilePath(fileId *string) *string {
	for _, item := range items {
		if item.Document.FileInfo.ID.UUID.String() == *fileId {
			item.Document.FileInfo.FileSystemPath = uploadHelper.BuildPath(fileId)
			return &item.Document.FileInfo.FileSystemPath
		}
	}
	return nil
}
