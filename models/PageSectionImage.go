package models

import (
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"github.com/uptrace/bun"
)

type PageSectionImage struct {
	bun.BaseModel  `bun:"page_section_images,alias:page_section_images"`
	ID             uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Description    string        `json:"description"`
	Order          uint          `bun:"document_type_image_order" json:"order"`
	DocumentTypeID uuid.NullUUID `bun:"type:uuid" json:"documentId" `
	FileInfo       *FileInfo     `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoID     uuid.NullUUID `bun:"type:uuid" json:"fileInfoId"`
}

type PageSectionImages []*PageSectionImage

func (i PageSectionImages) SetFileInfoID() {
	for _, item := range i {
		item.FileInfoID = item.FileInfo.ID
	}
}

func (i PageSectionImages) GetFileInfos() FileInfos {
	items := make(FileInfos, 0)
	for _, item := range i {
		items = append(items, item.FileInfo)
	}
	return items
}

func (i PageSectionImages) SetFilePath(fileID *string) *string {
	for _, item := range i {
		if item.FileInfo.ID.UUID.String() == *fileID {
			item.FileInfo.FileSystemPath = uploadHelper.BuildPath(fileID)
			return &item.FileInfo.FileSystemPath
		}
	}
	return nil
}
