package models

import (
	"github.com/pro-assistance/pro-assister/uploadHelper"

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

func (i DivisionImages) GetFileInfos() FileInfos {
	items := make(FileInfos, 0)
	for _, item := range i {
		items = append(items, item.FileInfo)
	}
	return items
}

func (i DivisionImages) SetFileInfoID() {
	for _, item := range i {
		item.FileInfoID = item.FileInfo.ID
	}
}

func (i DivisionImages) SetFilePath(fileID *string) *string {
	for _, item := range i {
		if item.FileInfo.ID.UUID.String() == *fileID {
			item.FileInfo.FileSystemPath = uploadHelper.BuildPath(fileID)
			return &item.FileInfo.FileSystemPath
		}
	}
	return nil
}
