package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DocumentType struct {
	bun.BaseModel        `bun:"document_types,alias:document_types"`
	ID                   uuid.UUID           `bun:"type:uuid,default:uuid_generate_v4()" json:"id,omitempty"`
	Name                 string              `json:"name,omitempty"`
	PublicDocumentTypeID uuid.NullUUID       `bun:"type:uuid,nullzero,default:NULL" json:"publicDocumentTypeId"`
	PublicDocumentType   *PublicDocumentType `bun:"rel:belongs-to" json:"publicDocumentType"`

	Documents          Documents   `bun:"rel:has-many" json:"documents"`
	DocumentsForDelete []uuid.UUID `bun:"-" json:"documentsForDelete"`

	DocumentTypeFields      DocumentTypeFields `bun:"rel:has-many" json:"documentFields"`
	DocumentFieldsForDelete []uuid.UUID        `bun:"-" json:"documentFieldsForDelete"`

	// SingleScan    bool      `json:"singleScan"`

	// Scan   *FileInfo     `bun:"rel:belongs-to" json:"scan,omitempty"`
	// ScanID uuid.NullUUID `bun:"type:uuid" json:"scanId"`

	// Scans          FileInfos `bun:"rel:has-many" json:"scans"`
	// ScansForDelete []string  `bun:"-" json:"scansForDelete"`
}

type DocumentTypes []*DocumentType

func (item *DocumentType) SetIdForChildren() {
	for i := range item.DocumentTypeFields {
		item.DocumentTypeFields[i].DocumentTypeID = item.ID
	}
	for i := range item.Documents {
		item.Documents[i].DocumentTypeID = item.ID
	}
}

func (items DocumentTypes) SetIdForChildren() {
	for i := range items {
		items[i].SetIdForChildren()
	}
}

func (item DocumentType) SetFilePath(fileID *string) *string {
	for i := range item.Documents {
		filePath := item.Documents[i].SetFilePath(fileID)
		if filePath != nil {
			return filePath
		}
	}
	return nil
}

func (items DocumentTypes) GetDocuments() Documents {
	itemsForGet := make(Documents, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.Documents...)
	}
	return itemsForGet
}

func (items DocumentTypes) GetDocumentsIdForDelete() []uuid.UUID {
	idPool := make([]uuid.UUID, 0)
	for _, item := range items {
		idPool = append(idPool, item.DocumentsForDelete...)
	}
	return idPool
}

// func (items DocumentTypes) GetFileInfos() FileInfos {
// 	itemsForGet := make(FileInfos, 0)
// 	for _, item := range items {
// 		itemsForGet = append(itemsForGet, item.Scans...)
// 		itemsForGet = append(itemsForGet, item.Scan)
// 	}
// 	return itemsForGet
// }

// func (items DocumentTypes) SetFileInfoID() {
// 	for _, item := range items {
// 		item.ScanID = item.Scan.ID
// 	}
// }
