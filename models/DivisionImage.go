package models

import (
	"github.com/pro-assistance/pro-assister/helpers/uploader"

	"github.com/google/uuid"
)

type DivisionImage struct {
	ID          uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Description string        `json:"description"`
	DivisionID  uuid.NullUUID `bun:"type:uuid" json:"divisionId" `
	FileInfo    *FileInfo     `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoID  uuid.NullUUID `bun:"type:uuid" json:"fileInfoId"`
}

type DivisionImages []*DivisionImage

func (items DivisionImages) GetFileInfos() FileInfos {
	itemsForGet := make(FileInfos, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.FileInfo)
	}
	return itemsForGet
}

func (items DivisionImages) SetFileInfoID() {
	for _, item := range items {
		item.FileInfoID = item.FileInfo.ID
	}
}

func (items DivisionImages) SetFilePath(fileID *string) *string {
	for i := range items {
		if items[i].FileInfo.ID.UUID.String() == *fileID {
			items[i].FileInfo.FileSystemPath = uploader.BuildPath(fileID)
			return &items[i].FileInfo.FileSystemPath
		}
	}
	return nil
}
