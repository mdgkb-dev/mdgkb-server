package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ResidencyDocumentType struct {
	bun.BaseModel  `bun:"residency_document_types,alias:residency_document_types"`
	ID             uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id,omitempty"`
	DocumentTypeID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"documentTypeId"`
	DocumentType   *DocumentType `bun:"rel:belongs-to" json:"documentType"`
}

type ResidencyDocumentTypes []*ResidencyDocumentType

func (item *ResidencyDocumentType) SetForeignKeys() {
	item.DocumentTypeID = item.DocumentType.ID
}

func (items ResidencyDocumentTypes) SetForeignKeys() {
	for i := range items {
		items[i].SetForeignKeys()
	}
}

func (item *ResidencyDocumentType) SetFilePath(fileID *string) *string {
	path := item.DocumentType.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}

func (items ResidencyDocumentTypes) SetFilePath(fileID *string) *string {
	for i := range items {
		path := items[i].SetFilePath(fileID)
		if path != nil {
			return path
		}
	}
	return nil
}

func (items ResidencyDocumentTypes) GetDocumentTypes() DocumentTypes {
	itemsForGet := make(DocumentTypes, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.DocumentType)
	}
	return itemsForGet
}
