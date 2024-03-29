package models

import (
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helpers/uploader"
	"github.com/uptrace/bun"
)

type FieldValueFile struct {
	bun.BaseModel `bun:"field_values_files,alias:field_values_files"`
	ID            uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	FieldValueID  uuid.UUID     `bun:"type:uuid" json:"fieldValueId"`
	FieldValue    *FieldValue   `bun:"rel:belongs-to" json:"fieldValue"`
	FileInfo      *FileInfo     `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoID    uuid.NullUUID `json:"fileInfoId"`
}

type FieldValuesFiles []*FieldValueFile

func (items FieldValuesFiles) SetFileInfoID() {
	for i := range items {
		items[i].FileInfoID = items[i].FileInfo.ID
	}
}

func (items FieldValuesFiles) GetFileInfos() FileInfos {
	itemsForGet := make(FileInfos, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.FileInfo)
	}
	return itemsForGet
}

func (item FieldValueFile) SetFilePath(fileID *string) *string {
	if item.FileInfo.ID.UUID.String() == *fileID {
		item.FileInfo.FileSystemPath = uploader.BuildPath(fileID)
		return &item.FileInfo.FileSystemPath
	}
	return nil
}

func (items FieldValuesFiles) SetFilePath(fileID *string) *string {
	for i := range items {
		filePath := items[i].SetFilePath(fileID)
		if filePath != nil {
			return filePath
		}
	}
	return nil
}
