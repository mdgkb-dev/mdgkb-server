package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type FormValueFile struct {
	bun.BaseModel `bun:"form_value_files,alias:form_value_files"`
	ID            uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	FormValue     *FormValue    `bun:"rel:belongs-to" json:"formValue"`
	FormValueID   uuid.NullUUID `bun:"type:uuid" json:"formValueId"`

	File   *FileInfo     `bun:"rel:belongs-to" json:"file"`
	FileID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"fileId"`
}

type FormValueFiles []*FormValueFile

func (items FormValueFiles) GetFiles() FileInfos {
	itemsForGet := make(FileInfos, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].File)
	}
	return itemsForGet
}

func (items FormValueFiles) SetForeignKeys() {
	for i := range items {
		items[i].FileID = items[i].File.ID
	}
}
