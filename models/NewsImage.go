package models

import (
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/uploadHelper"
)

type NewsImage struct {
	ID          uuid.UUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Description string        `json:"description"`
	NewsID      uuid.UUID     `bun:"type:uuid" json:"newsId" `
	FileInfo    *FileInfo     `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoID  uuid.NullUUID `bun:"type:uuid" json:"fileInfoId"`
}

type NewsImages []*NewsImage

func (i NewsImages) SetFileInfoID() {
	for _, item := range i {
		item.FileInfoID = item.FileInfo.ID
	}
}

func (i NewsImages) GetFileInfos() FileInfos {
	items := make(FileInfos, 0)
	for _, item := range i {
		items = append(items, item.FileInfo)
	}
	return items
}

func (i NewsImages) SetFilePath(fileID *string) *string {
	for _, item := range i {
		if item.FileInfo.ID.UUID.String() == *fileID {
			item.FileInfo.FileSystemPath = uploadHelper.BuildPath(fileID)
			return &item.FileInfo.FileSystemPath
		}
	}
	return nil
}
