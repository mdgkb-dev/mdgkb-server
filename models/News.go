package models

import (
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"github.com/uptrace/bun"
	"time"

	"github.com/google/uuid"
)

type News struct {
	bun.BaseModel        `bun:"news,alias:news"`
	ID                   uuid.UUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Status               string        `json:"status"`
	Title                string        `json:"title"`
	PreviewText          string        `json:"previewText"`
	Content              string        `json:"content"`
	Slug                 string        `json:"slug"`
	PublishedOn          time.Time     `json:"publishedOn"`
	Description          string        `json:"description"`
	Main                 bool          `json:"main"`
	SubMain              bool          `json:"subMain"`
	ArticleLink          string        `json:"articleLink"`
	IsArticle            bool          `json:"isArticle"`
	NewsImagesForDelete  []string      `bun:"-" json:"newsImagesForDelete"`
	NewsImagesNames      []string      `bun:"-" json:"newsImagesNames"`
	FileInfo             *FileInfo     `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoID           uuid.NullUUID `bun:"type:uuid" json:"fileInfoId"`
	MainImage            *FileInfo     `bun:"rel:belongs-to" json:"mainImage"`
	MainImageID          uuid.NullUUID `bun:"type:uuid" json:"mainImageId"`
	MainImageDescription string        `bun:"type:uuid" json:"mainImageDescription"`
	ViewsCount           int           `bun:"-" json:"viewsCount"`
	Event                *Event        `bun:"rel:belongs-to" json:"event"`
	EventID              uuid.NullUUID `bun:"type:uuid" json:"eventId"`

	NewsToCategories       NewsToCategories `bun:"rel:has-many" json:"newsToCategories"`
	NewsToTags             NewsToTags       `bun:"rel:has-many" json:"newsToTags"`
	NewsDoctors            NewsDoctors      `bun:"rel:has-many" json:"newsDoctors"`
	NewsDoctorsForDelete   []uuid.UUID      `bun:"-" json:"newsDoctorsForDelete"`
	NewsDivisions          NewsDivisions    `bun:"rel:has-many" json:"newsDivisions"`
	NewsDivisionsForDelete []uuid.UUID      `bun:"-" json:"newsDoctorsForDeleteForDelete"`
	NewsViews              NewsViews        `bun:"rel:has-many" json:"newsViews"`
	NewsLikes              NewsLikes        `bun:"rel:has-many" json:"newsLikes"`
	NewsComments           NewsComments     `bun:"rel:has-many" json:"newsComments"`
	NewsImages             NewsImages       `bun:"rel:has-many" json:"newsImages"`
}

func (item *News) SetFilePath(fileID *string) *string {
	if item.FileInfo.ID.UUID.String() == *fileID {
		item.FileInfo.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.FileInfo.FileSystemPath
	}
	if item.MainImage.ID.UUID.String() == *fileID {
		item.MainImage.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.MainImage.FileSystemPath
	}
	path := item.NewsImages.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}

func (item *News) SetForeignKeys() {
	item.FileInfoID = item.FileInfo.ID
	item.MainImageID = item.MainImage.ID
	if item.Event != nil {
		item.EventID = item.Event.ID
	}

}

func (item *News) SetIdForChildren() {
	for i := range item.NewsToCategories {
		item.NewsToCategories[i].NewsID = item.ID
	}
	for i := range item.NewsViews {
		item.NewsViews[i].NewsID = item.ID
	}
	for i := range item.NewsToTags {
		item.NewsToTags[i].NewsID = item.ID
	}
	for i := range item.NewsLikes {
		item.NewsLikes[i].NewsID = item.ID
	}
	for i := range item.NewsComments {
		item.NewsComments[i].NewsID = item.ID
	}
	for i := range item.NewsImages {
		item.NewsImages[i].NewsID = item.ID
	}
	for i := range item.NewsDoctors {
		item.NewsDoctors[i].NewsID = item.ID
	}
	for i := range item.NewsDivisions {
		item.NewsDivisions[i].NewsID = item.ID
	}
}

func (item *News) GetFileInfos() FileInfos {
	items := make(FileInfos, 0)
	items = append(items, item.FileInfo, item.MainImage)
	return items
}
