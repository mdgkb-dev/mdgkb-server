package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Certificate struct {
	bun.BaseModel `bun:"doctor_comments,alias:doctor_comments"`
	ID            uuid.UUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	DoctorID      uuid.NullUUID `bun:"type:uuid" json:"doctorId"`
	Doctor        *Doctor       `bun:"rel:belongs-to" json:"doctor"`
	Scan          *FileInfo     `bun:"rel:belongs-to" json:"scan"`
	ScanID        uuid.NullUUID `bun:"type:uuid" json:"scanId"`
	Description   string        `json:"description"`
}

type Certificates []*Certificate

func (items Certificates) GetFileInfos() FileInfos {
	itemsForGet := make(FileInfos, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.Scan)
	}
	return itemsForGet
}

func (items Certificates) SetForeignKeys() {
	for i := range items {
		items[i].ScanID = items[i].Scan.ID
	}
}
