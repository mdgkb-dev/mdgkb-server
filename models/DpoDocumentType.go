package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DpoDocumentType struct {
	bun.BaseModel  `bun:"dpo_document_types,alias:dpo_document_types"`
	ID             uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id,omitempty"`
	DocumentTypeID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"documentTypeId"`
	DocumentType   *PageSection  `bun:"rel:belongs-to" json:"documentType"`
}

type DpoDocumentTypes []*DpoDocumentType

func (item *DpoDocumentType) SetForeignKeys() {
	item.DocumentTypeID = item.DocumentType.ID
}

func (items DpoDocumentTypes) SetForeignKeys() {
	for i := range items {
		items[i].SetForeignKeys()
	}
}

func (item *DpoDocumentType) SetFilePath(fileID *string) *string {
	path := item.DocumentType.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}

func (items DpoDocumentTypes) SetFilePath(fileID *string) *string {
	for i := range items {
		path := items[i].SetFilePath(fileID)
		if path != nil {
			return path
		}
	}
	return nil
}

func (items DpoDocumentTypes) GetDocumentTypes() PageSections {
	itemsForGet := make(PageSections, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.DocumentType)
	}
	return itemsForGet
}
