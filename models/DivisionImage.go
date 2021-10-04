package models

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/helpers/uploadHelper"
)

type DivisionImage struct {
	ID          uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Description string    `json:"description"`
	DivisionId  uuid.UUID `bun:"type:uuid" json:"divisionId" `
	FileInfo    *FileInfo `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoId  uuid.NullUUID `bun:"type:uuid" json:"fileInfoId"`
}

type DivisionImages []*DivisionImage

func (i DivisionImages) GetFileInfos() FileInfos {
	items := make(FileInfos, 0)
	for _, item := range i {
		items = append(items, item.FileInfo)
	}
	return items
}

func (i DivisionImages) SetFileInfoID()  {
	for _, item := range i {
		item.FileInfoId = item.FileInfo.ID
	}
}

func (i DivisionImages) SetFilePath(fileId *string) *string {
	for _, item := range i {
		if item.FileInfo.ID.UUID.String() == *fileId {
			item.FileInfo.FileSystemPath = uploadHelper.BuildPath(fileId)
			return &item.FileInfo.FileSystemPath
		}
	}
	return nil
}
