package models

import (
	"github.com/google/uuid"
)

type FileInfo struct {
	ID             uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	OriginalName   string    `json:"originalName"`
	FileSystemPath string    `json:"fileSystemPath"`
}

type FileInfos []*FileInfo