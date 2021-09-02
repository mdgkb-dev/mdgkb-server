package models

import (
	"github.com/google/uuid"
)

type DivisionImage struct {
	ID          uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Description string    `json:"description"`
	DivisionId  uuid.UUID `bun:"type:uuid" json:"divisionId" `
	FileInfo    *FileInfo `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoId  uuid.UUID `bun:"type:uuid" json:"fileInfoId"`
}
