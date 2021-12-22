package models

import (
	"mdgkb/mdgkb-server/helpers/uploadHelper"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Partner struct {
	bun.BaseModel `bun:"partners,alias:partners"`
	ID            uuid.UUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Link          string        `json:"link"`
	PartnerType   *PartnerType  `bun:"rel:belongs-to" json:"partnerType"`
	PartnerTypeID uuid.NullUUID `bun:"type:uuid" json:"partnerTypeId"`
	Image         *FileInfo     `bun:"rel:belongs-to" json:"image"`
	ImageID       uuid.NullUUID `bun:"type:uuid" json:"imageId"`
}

type Partners []*Partner

func (item *Partner) SetFilePath(fileID *string) *string {
	if item.Image.ID.UUID.String() == *fileID {
		item.Image.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.Image.FileSystemPath
	}
	return nil
}

func (item *Partner) GetFileInfos() FileInfos {
	items := make(FileInfos, 0)
	items = append(items, item.Image)
	return items
}

func (item *Partner) SetForeignKeys() {
	if item.Image != nil {
		item.ImageID = item.Image.ID
	}
	item.PartnerTypeID = item.PartnerType.ID
}
