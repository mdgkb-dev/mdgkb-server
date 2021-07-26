package models

import (
	"github.com/google/uuid"
)

type FileInfo struct {
	ID           uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	FilenameDisk string    `json:"filenameDisk"`
	OriginalName string    `json:"originalName"`
}
