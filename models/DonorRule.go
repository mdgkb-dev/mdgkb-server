package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"github.com/pro-assistance/pro-assister/uploadHelper"
)

type DonorRule struct {
	bun.BaseModel   `bun:"donor_rules,alias:donor_rules"`
	ID              uuid.UUID       `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name            string          `json:"name"`
	Order           int             `bun:"donor_rule_order" json:"order"`
	Image           *FileInfo       `bun:"rel:belongs-to" json:"image"`
	ImageID         uuid.NullUUID   `bun:"type:uuid" json:"imageId"`
	DonorRulesUsers DonorRulesUsers `bun:"rel:has-many" json:"donorRulesUsers"`
}

type DonorRules []*DonorRule

func (items DonorRules) SetForeignKeys() {
	for i := range items {
		items[i].ImageID = items[i].Image.ID
	}
}

func (items DonorRules) GetImages() FileInfos {
	itemsForGet := make(FileInfos, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.Image)
	}
	return itemsForGet
}

func (item *DonorRule) SetFilePath(fileID string) *string {
	if item.Image.ID.UUID.String() == fileID {
		item.Image.FileSystemPath = uploadHelper.BuildPath(&fileID)
		return &item.Image.FileSystemPath
	}
	return nil
}

func (items DonorRules) SetFilePath(fileID string) *string {
	var filePath *string
	for i := range items {
		path := items[i].SetFilePath(fileID)
		if path != nil {
			filePath = path
			break
		}
	}
	return filePath
}
