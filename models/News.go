package models

import (
	"github.com/google/uuid"
	"time"
)

type News struct {
	ID                     uuid.UUID             `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Status                 string                `json:"status"`
	Title                  string                `json:"title"`
	PreviewText            string                `json:"preview_text"`
	Content                string                `json:"content"`
	Slug                   string                `json:"slug"`
	PublishedOn            time.Time             `json:"publishedOn"`
	Description            string                `json:"description"`
	Categories             []Category            `bun:"m2m:news_to_categories" json:"categories"`
	Tags                   []Tag                 `bun:"m2m:news_to_tags" json:"tags"`
	NewsLikes              []*NewsLike           `bun:"rel:has-many" json:"newsLikes"`
	NewsComments           []*NewsComment        `bun:"rel:has-many" json:"newsComments"`
	PreviewThumbnailFile   *PreviewThumbnailFile `bun:"rel:belongs-to" json:"previewThumbnailFile"`
	PreviewThumbnailFileId uuid.UUID             `bun:"type:uuid" json:"previewThumbnailFileId"`
}

type NewsToCategory struct {
	ID         uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	NewsId     uuid.UUID `bun:"type:uuid"`
	News       *News     `bun:"rel:has-one"`
	CategoryID uuid.UUID `bun:"type:uuid"`
	Category   *Category `bun:"rel:has-one"`
}

type NewsToTag struct {
	ID     uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	NewsId uuid.UUID `bun:"type:uuid"`
	News   *News     `bun:"rel:has-one"`
	TagId  uuid.UUID `bun:"type:uuid"`
	Tag    *Tag      `bun:"rel:has-one"`
}
