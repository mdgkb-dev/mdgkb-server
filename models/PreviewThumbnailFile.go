package models

import (
	"github.com/google/uuid"
)

type PreviewThumbnailFile struct {
	ID           uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	FilenameDisk string    `json:"filenameDisk"`
}
