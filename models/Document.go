package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Document struct {
	bun.BaseModel `bun:"documents,alias:documents" json:"bun_base_model"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" json:"id,omitempty"`
	Name          string    `json:"name" json:"name,omitempty"`
	FileInfo      *FileInfo `bun:"rel:belongs-to" json:"fileInfo" json:"file_info,omitempty"`
	FileInfoId    uuid.UUID `bun:"type:uuid" json:"fileInfoId"`
}

type Documents []*Document

func (i Documents) GetFileInfos() FileInfos {
	items := make(FileInfos, 0)
	for _, item := range i {
		items = append(items, item.FileInfo)
	}
	return items
}

func (i Documents) SetFileInfoID()  {
	for _, item := range i {
		item.FileInfoId = item.FileInfo.ID
	}
}