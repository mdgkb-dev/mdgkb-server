package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/pro-assistance/pro-assister/uploadHelper"
)

type Banner struct {
	bun.BaseModel `bun:"banners,alias:banners"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Link          string        `json:"link"`
	ListNumber    int           `bun:"type:integer" json:"listNumber"`
	FileInfo      *FileInfo     `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoId    uuid.UUID     `bun:"type:uuid" json:"fileInfoId"`
}

type Banners []*Banner

func (item *Banner) SetFilePath(fileID *string) *string {
	if item.FileInfo.ID.UUID.String() == *fileID {
		item.FileInfo.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.FileInfo.FileSystemPath
	}
	return nil
}

func (item *Banner) SetForeignKeys() {
	if item.FileInfo != nil {
		item.FileInfoId = item.FileInfo.ID.UUID
	}
}
