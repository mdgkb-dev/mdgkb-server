package models

import (
	"fmt"
	"github.com/pro-assistance/pro-assister/uploadHelper"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Document struct {
	bun.BaseModel `bun:"documents,alias:documents"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	Name          string    `json:"name"`

	DocumentType   *DocumentType `bun:"rel:belongs-to" json:"documentType"`
	DocumentTypeID uuid.UUID     `bun:"type:uuid" json:"documentTypeId"`

	DocumentsScans          DocumentsScans      `bun:"rel:has-many" json:"documentsScans"`
	DocumentsScansForDelete []uuid.UUID         `bun:"-" json:"documentsScansForDelete"`
	DocumentFieldsValues    DocumentFieldValues `bun:"rel:has-many" json:"documentFields"`
}

type Documents []*Document

func (item *Document) SetIdForChildren() {
	for i := range item.DocumentFieldsValues {
		item.DocumentFieldsValues[i].DocumentID = item.ID
	}
	for i := range item.DocumentsScans {
		item.DocumentsScans[i].DocumentID = item.ID
	}
}

func (items Documents) SetIdForChildren() {
	for i := range items {
		items[i].SetIdForChildren()
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

func (items Documents) GetDocumentsScansIdForDelete() []uuid.UUID {
	idPool := make([]uuid.UUID, 0)
	for _, item := range items {
		idPool = append(idPool, item.DocumentsScansForDelete...)
	}
	return idPool
}

func (item *Document) SetFilePath(fileID *string) *string {
	for i, documentScan := range item.DocumentsScans {
		fmt.Println(documentScan.Scan.ID.UUID.String())
		fmt.Println(*fileID)
		fmt.Println(documentScan.Scan.ID.UUID.String() == *fileID)
		fmt.Println(i)
		if documentScan.Scan.ID.UUID.String() == *fileID {
			item.DocumentsScans[i].Scan.FileSystemPath = uploadHelper.BuildPath(fileID)
			return &item.DocumentsScans[i].Scan.FileSystemPath
		}
	}
	fmt.Println("len", len(item.DocumentsScans))
	return nil
}

func (items Documents) SetFilePath(fileID *string) *string {
	for i := range items {
		items[i].SetFilePath(fileID)
	}
	return nil
}
