package models

import (
	"github.com/google/uuid"
)

type NormativeDocument struct {
	ID                      uuid.UUID              `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	Name                    string                 `json:"name"`
	NormativeDocumentType   *NormativeDocumentType `bun:"rel:belongs-to" json:"type"`
	NormativeDocumentTypeId uuid.UUID              `bun:"type:uuid"`
	FileInfo                *FileInfo              `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoId              uuid.UUID              `bun:"type:uuid"`
}
