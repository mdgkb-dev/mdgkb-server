package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Document struct {
	bun.BaseModel `bun:"documents,alias:documents" json:"bun_base_model"`
	ID            uuid.UUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" json:"id,omitempty"`
	Name          string        `json:"name" json:"name,omitempty"`
	FileInfo      *FileInfo     `bun:"rel:belongs-to" json:"fileInfo" json:"file_info,omitempty"`
	FileInfoId    uuid.NullUUID `bun:"type:uuid" json:"fileInfoId"`

	Scans          FileInfos `bun:"rel:has-many" json:"scans"`
	ScansForDelete []string  `bun:"-" json:"scansForDelete"`

	DocumentFields          DocumentFields `bun:"rel:has-many" json:"documentFields"`
	DocumentFieldsForDelete []uuid.UUID    `bun:"-" json:"documentFieldsForDelete"`
}

type Documents []*Document

func (item *Document) SetIdForChildren() {
	for i := range item.DocumentFields {
		item.DocumentFields[i].DocumentID = item.ID
	}
}

func (items Documents) SetIdForChildren() {
	for i := range items {
		items[i].SetIdForChildren()
	}
}

func (items Documents) GetFileInfos() FileInfos {
	itemsForGet := make(FileInfos, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.FileInfo)
	}
	return itemsForGet
}

func (items Documents) SetFileInfoID() {
	for _, item := range items {
		item.FileInfoId = item.FileInfo.ID
	}
}
