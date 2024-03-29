package models

import (
	"github.com/pro-assistance/pro-assister/helpers/uploader"

	"github.com/google/uuid"
)

type PageImage struct {
	ID          uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Description string        `json:"description"`
	PageID      uuid.NullUUID `bun:"type:uuid" json:"pageID" `
	FileInfo    *FileInfo     `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoID  uuid.NullUUID `bun:"type:uuid" json:"fileInfoId"`
}

type PageImages []*PageImage

func (i PageImages) GetFileInfos() FileInfos {
	items := make(FileInfos, 0)
	for _, item := range i {
		items = append(items, item.FileInfo)
	}
	return items
}

func (i PageImages) SetFileInfoID() {
	for _, item := range i {
		item.FileInfoID = item.FileInfo.ID
	}
}

func (i PageImages) SetFilePath(fileID *string) *string {
	for _, item := range i {
		if item.FileInfo.ID.UUID.String() == *fileID {
			item.FileInfo.FileSystemPath = uploader.BuildPath(fileID)
			return &item.FileInfo.FileSystemPath
		}
	}
	return nil
}
