package models

import (
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helpers/uploader"
	"github.com/uptrace/bun"
)

type FileInfo struct {
	bun.BaseModel  `bun:"file_infos,alias:file_infos"`
	ID             uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	OriginalName   string        `json:"originalName"`
	FileSystemPath string        `json:"fileSystemPath"`
}

// TODO: FileInfo2 временное решение в связи с багом bun

type FileInfo2 struct {
	bun.BaseModel  `bun:"file_infos,alias:file_infos"`
	ID             uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	OriginalName   string        `json:"originalName"`
	FileSystemPath string        `json:"fileSystemPath"`
}

type FileInfos []*FileInfo

func (item FileInfo) GetOriginalName() string {
	return item.OriginalName
}

func (item FileInfo) GetFullPath() string {
	return item.FileSystemPath
}
func (items FileInfos) GetPathsAndNames() (paths []string, names []string) {
	for _, item := range items {
		paths = append(paths, item.FileSystemPath)
		names = append(names, item.OriginalName)
	}
	return paths, names
}

func (item *FileInfo) SetFilePath(fileID *string) *string {
	if item.ID.UUID.String() == *fileID {
		item.FileSystemPath = uploader.BuildPath(fileID)
		return &item.FileSystemPath
	}
	return nil
}
