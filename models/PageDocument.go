package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PageDocument struct {
	bun.BaseModel `bun:"pages_documents,alias:pages_documents"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`

	Page   *Page     `bun:"rel:belongs-to" json:"page"`
	PageID uuid.UUID `bun:"type:uuid" json:"pageId"`

	Document   *Document `bun:"rel:belongs-to" json:"document"`
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

func (items PageDocuments) SetForeignKeys() {
	for i := range items {
		items[i].DocumentID = items[i].Document.ID
	}
}

func (items PageDocuments) SetFilePath(fileID *string) *string {
	for _ = range items {
		//if item.Document.Scan.ID.UUID.String() == *fileID {
		//	item.Document.Scan.FileSystemPath = uploadHelper.BuildPath(fileID)
		//	return &item.Document.Scan.FileSystemPath
		//}
	}
	return nil
}
