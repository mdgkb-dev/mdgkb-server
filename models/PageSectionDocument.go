package models

import (
	"github.com/pro-assistance/pro-assister/uploadHelper"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PageSectionDocument struct {
	bun.BaseModel  `bun:"page_section_documents,alias:page_section_documents"`
	ID             uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name           string        `json:"name"`
	Order          uint          `bun:"item_order" json:"order"`
	PageSection    *PageSection  `bun:"rel:belongs-to" json:"pageSection"`
	PageSectionID  uuid.NullUUID `bun:"type:uuid" json:"pageSectionId"`
	DownloadToFile bool          `json:"downloadToFile"`

	Scan   *FileInfo     `bun:"rel:belongs-to" json:"scan"`
	ScanID uuid.NullUUID `bun:"type:uuid" json:"scanId"`
}

type PageSectionDocuments []*PageSectionDocument

func (item *PageSectionDocument) SetForeignKeys() {
	item.ScanID = item.Scan.ID
}

func (items PageSectionDocuments) SetForeignKeys() {
	for i := range items {
		items[i].SetForeignKeys()
	}
}

func (item *PageSectionDocument) SetFilePath(fileID *string) *string {
	if item.Scan.ID.UUID.String() == *fileID {
		item.Scan.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.Scan.FileSystemPath
	}
	return nil
}

func (items PageSectionDocuments) SetFilePath(fileID *string) *string {
	for i := range items {
		items[i].SetFilePath(fileID)
	}
	return nil
}

func (items PageSectionDocuments) GetScans() FileInfos {
	itemsForGet := make(FileInfos, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.Scan)
	}
	return itemsForGet
}
