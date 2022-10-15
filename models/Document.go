package models

import (
	"github.com/pro-assistance/pro-assister/uploadHelper"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Document struct {
	bun.BaseModel           `bun:"documents,alias:documents"`
	ID                      uuid.UUID           `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name                    string              `json:"name"`
	Order                   uint                `bun:"document_order" json:"order"`
	DocumentType            *DocumentType       `bun:"rel:belongs-to" json:"documentType"`
	DocumentTypeID          uuid.NullUUID       `bun:"type:uuid" json:"documentTypeId"`
	DownloadToFile          bool                `json:"downloadToFile"`
	DocumentsScans          DocumentsScans      `bun:"rel:has-many" json:"documentsScans"`
	DocumentsScansForDelete []uuid.UUID         `bun:"-" json:"documentsScansForDelete"`
	DocumentFieldsValues    DocumentFieldValues `bun:"rel:has-many" json:"documentFields"`
}

type Documents []*Document

func (item *Document) SetIDForChildren() {
	for i := range item.DocumentFieldsValues {
		item.DocumentFieldsValues[i].DocumentID = item.ID
	}
	for i := range item.DocumentsScans {
		item.DocumentsScans[i].DocumentID = item.ID
	}
}

func (items Documents) SetIDForChildren() {
	for i := range items {
		items[i].SetIDForChildren()
	}
}

func (items Documents) GetDocumentsScans() DocumentsScans {
	itemsForGet := make(DocumentsScans, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.DocumentsScans...)
		//itemsForGet = append(itemsForGet, item.Scan)
	}
	return itemsForGet
}

func (items Documents) GetDocumentsScansIDForDelete() []uuid.UUID {
	idPool := make([]uuid.UUID, 0)
	for _, item := range items {
		idPool = append(idPool, item.DocumentsScansForDelete...)
	}
	return idPool
}

func (item *Document) SetFilePath(fileID *string) *string {
	for i, documentScan := range item.DocumentsScans {
		if documentScan.Scan.ID.UUID.String() == *fileID {
			item.DocumentsScans[i].Scan.FileSystemPath = uploadHelper.BuildPath(fileID)
			return &item.DocumentsScans[i].Scan.FileSystemPath
		}
	}
	return nil
}

func (items Documents) SetFilePath(fileID *string) *string {
	for i := range items {
		items[i].SetFilePath(fileID)
	}
	return nil
}
