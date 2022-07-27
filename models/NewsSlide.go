package models

import (
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"github.com/uptrace/bun"
)

type NewsSlide struct {
	bun.BaseModel             `bun:"news_slides,alias:news_slides"`
	ID                        uuid.UUID        `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Title                     string           `json:"title"`
	Content                   string           `json:"content"`
	Color                     string           `json:"color"`
	Order                     uint             `bun:"news_slide_order" json:"order"`
	DesktopImg                *FileInfo        `bun:"rel:belongs-to" json:"desktopImg"`
	DesktopImgID              uuid.NullUUID    `bun:"type:uuid" json:"desktopImgId"`
	LaptopImg                 *FileInfo        `bun:"rel:belongs-to" json:"laptopImg"`
	LaptopImgID               uuid.NullUUID    `bun:"type:uuid" json:"laptopImgId"`
	TabletImg                 *FileInfo        `bun:"rel:belongs-to" json:"tabletImg"`
	TabletImgID               uuid.NullUUID    `bun:"type:uuid" json:"tabletImgId"`
	MobileImg                 *FileInfo        `bun:"rel:belongs-to" json:"mobileImg"`
	MobileImgID               uuid.NullUUID    `bun:"type:uuid" json:"mobileImgId"`
	NewsSlideButtons          NewsSlideButtons `bun:"rel:has-many" json:"newsSlideButtons"`
	NewsSlideButtonsForDelete []string         `bun:"-" json:"newsSlideButtonsForDelete"`
}

type NewsSlides []*NewsSlide

func (item *NewsSlide) SetFilePath(fileID *string) *string {
	if item.DesktopImg.ID.UUID.String() == *fileID {
		item.DesktopImg.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.DesktopImg.FileSystemPath
	}
	if item.LaptopImg.ID.UUID.String() == *fileID {
		item.LaptopImg.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.LaptopImg.FileSystemPath
	}
	if item.TabletImg.ID.UUID.String() == *fileID {
		item.TabletImg.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.TabletImg.FileSystemPath
	}
	if item.MobileImg.ID.UUID.String() == *fileID {
		item.MobileImg.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.MobileImg.FileSystemPath
	}
	return nil
}

func (item *NewsSlide) SetIDForChildren() {
	if len(item.NewsSlideButtons) == 0 {
		return
	}
	for i := range item.NewsSlideButtons {
		item.NewsSlideButtons[i].NewsSlideID = item.ID
	}
}

func (item *NewsSlide) GetFileInfos() FileInfos {
	items := make(FileInfos, 0)
	items = append(items, item.DesktopImg, item.LaptopImg, item.TabletImg, item.MobileImg)
	return items
}

func (item *NewsSlide) SetForeignKeys() {
	item.DesktopImgID = item.DesktopImg.ID
	item.LaptopImgID = item.LaptopImg.ID
	item.TabletImgID = item.TabletImg.ID
	item.MobileImgID = item.MobileImg.ID
}
