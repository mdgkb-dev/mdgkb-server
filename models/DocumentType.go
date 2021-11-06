package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DocumentType struct {
	bun.BaseModel `bun:"document_types,alias:document_types"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" json:"id,omitempty"`
	Name          string    `json:"name" json:"name,omitempty"`
	SingleScan    bool      `json:"singleScan"`

	Scan   *FileInfo     `bun:"rel:belongs-to" json:"scan,omitempty"`
	ScanID uuid.NullUUID `bun:"type:uuid" json:"scanId"`

	Scans          FileInfos `bun:"rel:has-many" json:"scans"`
	ScansForDelete []string  `bun:"-" json:"scansForDelete"`

	DocumentTypeFields      DocumentTypeFields `bun:"rel:has-many" json:"documentFields"`
	DocumentFieldsForDelete []uuid.UUID        `bun:"-" json:"documentFieldsForDelete"`
}

type DocumentsTypes []*DocumentType

func (item *DocumentType) SetIdForChildren() {
	for i := range item.DocumentTypeFields {
		item.DocumentTypeFields[i].DocumentTypeID = item.ID
	}
}

func (items DocumentsTypes) SetIdForChildren() {
	for i := range items {
		items[i].SetIdForChildren()
	}
}

func (items DocumentsTypes) GetFileInfos() FileInfos {
	itemsForGet := make(FileInfos, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.Scans...)
		itemsForGet = append(itemsForGet, item.Scan)
	}
	return itemsForGet
}

func (items DocumentsTypes) SetFileInfoID() {
	for _, item := range items {
		item.ScanID = item.Scan.ID
	}
}
