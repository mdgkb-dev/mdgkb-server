package models

import (
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helpers/uploader"
	"github.com/uptrace/bun"
)

type NewsImage struct {
	bun.BaseModel `bun:"news_images,alias:news_images"`
	ID            uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Description   string        `json:"description"`
	Order         uint          `bun:"news_image_order" json:"order"`
	NewsID        uuid.NullUUID `bun:"type:uuid" json:"newsId" `
	FileInfo      *FileInfo     `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoID    uuid.NullUUID `bun:"type:uuid" json:"fileInfoId"`
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
			item.FileInfo.FileSystemPath = uploader.BuildPath(fileID)
			return &item.FileInfo.FileSystemPath
		}
	}
	return nil
}
