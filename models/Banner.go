package models

import (
	"github.com/google/uuid"
)

type Banner struct {
	ID         uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name       string    `json:"name"`
	Link       string    `json:"link"`
	ListNumber int       `bun:"type:integer" json:"listNumber"`
	FileInfo   *FileInfo `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoId uuid.UUID `bun:"type:uuid" json:"fileInfoId"`
}
