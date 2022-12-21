package models

import (
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"github.com/uptrace/bun"
)

type DishSample struct {
	bun.BaseModel `bun:"dishes_samples,alias:dish_samples"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Price         uint          `json:"price"`
	Caloric       uint          `json:"caloric"`
	Weight        uint          `json:"weight"`
	DishesGroup   *DishesGroup  `bun:"rel:belongs-to" json:"dishesGroup"`
	DishesGroupID uuid.NullUUID `bun:"type:uuid"  json:"dishesGroupId"`

	Image   *FileInfo     `bun:"rel:belongs-to" json:"image"`
	ImageID uuid.NullUUID `bun:"type:uuid" json:"imageId"`
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
