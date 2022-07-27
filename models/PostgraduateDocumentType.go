package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PostgraduateDocumentType struct {
	bun.BaseModel  `bun:"postgraduate_document_types,alias:postgraduate_document_types"`
	ID             uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id,omitempty"`
	DocumentTypeID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"documentTypeId"`
	DocumentType   *DocumentType `bun:"rel:belongs-to" json:"documentType"`
}

type PostgraduateDocumentTypes []*PostgraduateDocumentType

func (item *PostgraduateDocumentType) SetForeignKeys() {
	item.DocumentTypeID = item.DocumentType.ID
}

func (items PostgraduateDocumentTypes) SetForeignKeys() {
	for i := range items {
		items[i].SetForeignKeys()
	}
}

func (item *PostgraduateDocumentType) SetFilePath(fileID *string) *string {
	path := item.DocumentType.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}

func (items PostgraduateDocumentTypes) SetFilePath(fileID *string) *string {
	for i := range items {
		path := items[i].SetFilePath(fileID)
		if path != nil {
			return path
		}
	}
	return nil
}

func (items PostgraduateDocumentTypes) GetDocumentTypes() DocumentTypes {
	itemsForGet := make(DocumentTypes, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.DocumentType)
	}
	return itemsForGet
}
