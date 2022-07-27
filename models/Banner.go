package models

import (
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"github.com/uptrace/bun"
)

type Banner struct {
	bun.BaseModel `bun:"banners,alias:banners"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Link          string        `json:"link"`
	ListNumber    int           `bun:"type:integer" json:"listNumber"`
	FileInfo      *FileInfo     `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoID    uuid.UUID     `bun:"type:uuid" json:"fileInfoId"`
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
		item.FileInfoID = item.FileInfo.ID.UUID
	}
}
