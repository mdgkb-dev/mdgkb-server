package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DocumentScan struct {
	bun.BaseModel `bun:"documents_scans,alias:documents_scans"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" json:"id,omitempty"`

	Document   *Document `bun:"rel:belongs-to" json:"document"`
	DocumentID uuid.UUID `bun:"type:uuid" json:"documentId"`

	Scan   *FileInfo     `bun:"rel:belongs-to" json:"scan"`
	ScanID uuid.NullUUID `bun:"type:uuid" json:"scanId"`
}

type DocumentsScans []*DocumentScan

func (items DocumentsScans) SetForeignKeys() {
	for i := range items {
		items[i].ScanID = items[i].Scan.ID
	}
}

//func (item *Document) SetIdForChildren() {
//	for i := range item.DocumentFieldsValues {
//		item.DocumentFieldsValues[i].DocumentID = item.ID
//	}
//}
//
//func (items Documents) SetIdForChildren() {
//	for i := range items {
//		items[i].SetIdForChildren()
//	}
//}

func (items DocumentsScans) GetFileInfos() FileInfos {
	itemsForGet := make(FileInfos, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.Scan)
	}
	return itemsForGet
}

//
//func (item *Document) SetFilePath(fileID *string) *string {
//	for i, scan := range item.Scans {
//		if scan.ID.UUID.String() == *fileID {
//			item.Scans[i].FileSystemPath = uploadHelper.BuildPath(fileID)
//			return &item.Scans[i].FileSystemPath
//		}
//	}
//	return nil
//}
//
//func (items Documents) SetFilePath(fileID *string) *string {
//	for i := range items {
//		items[i].SetFilePath(fileID)
//	}
//	return nil
//}
