package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"github.com/uptrace/bun"
)

type DishSample struct {
	bun.BaseModel `bun:"dishes_samples,alias:dishes_samples"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Price         uint          `json:"price"`
	Caloric       uint          `json:"caloric"`
	Weight        uint          `json:"weight"`
	Quantity      int           `json:"quantity"`
	DishesGroup   *DishesGroup  `bun:"rel:belongs-to" json:"dishesGroup"`
	DishesGroupID uuid.NullUUID `bun:"type:uuid"  json:"dishesGroupId"`
	Order         uint8         `bun:"item_order" json:"order"`
	Image         *FileInfo     `bun:"rel:belongs-to" json:"image"`
	ImageID       uuid.NullUUID `bun:"type:uuid" json:"imageId"`
	UpdatedAt     time.Time     `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`
}

type DishSamples []*DishSample

func (item *DishSample) SetFilePath(fileID *string) *string {
	if item.Image.ID.UUID.String() == *fileID {
		item.Image.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.Image.FileSystemPath
	}
	return nil
}

func (item *DishSample) SetForeignKeys() {
	item.ImageID = item.Image.ID
}
