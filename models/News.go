package models

import (
	"time"

	"github.com/pro-assistance/pro-assister/helpers/uploader"
	"github.com/uptrace/bun"

	"github.com/google/uuid"
)

type News struct {
	bun.BaseModel        `bun:"news,select:news_view,alias:news_view"`
	ID                   uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Status               string        `json:"status,omitempty"`
	Title                string        `json:"title,omitempty"`
	PreviewText          string        `json:"previewText,omitempty"`
	Content              string        `json:"content,omitempty"`
	Slug                 string        `json:"slug,omitempty"`
	PublishedOn          time.Time     `json:"publishedOn,omitempty"`
	CreatedAt            time.Time     `json:"createdAt,omitempty"`
	Description          string        `json:"description,omitempty"`
	Main                 bool          `json:"main,omitempty"`
	SubMain              bool          `json:"subMain,omitempty"`
	ArticleLink          string        `json:"articleLink,omitempty"`
	IsArticle            bool          `json:"isArticle,omitempty"`
	NewsImagesForDelete  []uuid.UUID   `bun:"-" json:"newsImagesForDelete,omitempty"`
	NewsImagesNames      []string      `bun:"-" json:"newsImagesNames,omitempty"`
	PreviewImage         *FileInfo     `bun:"rel:belongs-to" json:"previewImage,omitempty"`
	PreviewImageID       uuid.NullUUID `bun:"type:uuid" json:"previewImageId,omitempty"`
	MainImage            *FileInfo     `bun:"rel:belongs-to" json:"mainImage,omitempty"`
	MainImageID          uuid.NullUUID `bun:"type:uuid" json:"mainImageId,omitempty"`
	MainImageDescription string        `bun:"type:uuid" json:"mainImageDescription,omitempty"`
	ViewsCount           int           `bun:",scanonly" json:"viewsCount,omitempty"`
	Event                *Event        `bun:"rel:belongs-to" json:"event,omitempty"`
	EventID              uuid.NullUUID `bun:"type:uuid" json:"eventId,omitempty"`
	IsDraft              bool          `json:"isDraft,omitempty"`

	NewsToCategories       NewsToCategories `bun:"rel:has-many" json:"newsToCategories,omitempty"`
	NewsToTags             NewsToTags       `bun:"rel:has-many" json:"newsToTags,omitempty"`
	NewsToTagsForDelete    []string         `bun:"-" json:"newsToTagsForDelete,omitempty"`
	NewsDoctors            NewsDoctors      `bun:"rel:has-many" json:"newsDoctors,omitempty"`
	NewsDoctorsForDelete   []uuid.UUID      `bun:"-" json:"newsDoctorsForDelete,omitempty"`
	NewsDivisions          NewsDivisions    `bun:"rel:has-many" json:"newsDivisions,omitempty"`
	NewsDivisionsForDelete []uuid.UUID      `bun:"-" json:"newsDoctorsForDeleteForDelete,omitempty"`
	NewsViews              NewsViews        `bun:"rel:has-many" json:"newsViews,omitempty"`
	NewsLikes              NewsLikes        `bun:"rel:has-many" json:"newsLikes,omitempty"`
	Comments               Comments         `bun:"rel:has-many,join:id=item_id" json:"comments"`
	NewsImages             NewsImages       `bun:"rel:has-many" json:"newsImages,omitempty"`
	TagCount               uint             `bun:",scanonly" json:"tagCount,omitempty"`
}

func (item *News) SetFilePath(fileID *string) *string {
	if item.PreviewImage.ID.UUID.String() == *fileID {
		item.PreviewImage.FileSystemPath = uploader.BuildPath(fileID)
		return &item.PreviewImage.FileSystemPath
	}
	if item.MainImage.ID.UUID.String() == *fileID {
		item.MainImage.FileSystemPath = uploader.BuildPath(fileID)
		return &item.MainImage.FileSystemPath
	}
	path := item.NewsImages.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}

func (item *News) SetForeignKeys() {
	item.PreviewImageID = item.PreviewImage.ID
	item.MainImageID = item.MainImage.ID
	if item.Event != nil {
		item.EventID = item.Event.ID
	}
}

func (item *News) SetIDForChildren() {
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
	// for i := range item.NewsComments {
	// 	item.NewsComments[i].NewsID = item.ID
	// }
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
	items = append(items, item.PreviewImage, item.MainImage)
	return items
}
