package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type EducationDocumentType struct {
	bun.BaseModel          `bun:"education_document_types,alias:education_document_types"`
	ID                     uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id,omitempty"`
	Order                  int           `bun:"education_document_type_order" json:"order"`
	DocumentTypeID         uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"documentTypeId"`
	DocumentType           *DocumentType `bun:"rel:belongs-to" json:"documentType"`
	DocumentTypes          DocumentTypes `bun:"rel:has-many" json:"documentTypes"`
	DocumentTypesForDelete []uuid.UUID   `bun:"-" json:"documentTypesForDelete"`
}

type EducationDocumentTypes []*EducationDocumentType

func (item *EducationDocumentType) SetForeignKeys() {
	item.DocumentTypeID = item.DocumentType.ID
}

func (items EducationDocumentTypes) SetForeignKeys() {
	for i := range items {
		items[i].SetForeignKeys()
	}
}

func (item *EducationDocumentType) SetFilePath(fileID *string) *string {
	path := item.DocumentType.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}

func (items EducationDocumentTypes) SetFilePath(fileID *string) *string {
	for i := range items {
		path := items[i].SetFilePath(fileID)
		if path != nil {
			return path
		}
	}
	return nil
}

func (items EducationDocumentTypes) GetDocumentTypes() DocumentTypes {
	itemsForGet := make(DocumentTypes, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.DocumentType)
	}
	return itemsForGet
}
