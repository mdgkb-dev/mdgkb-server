package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID                   uuid.UUID             `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Status               string                `json:"status"`
	Email                string                `json:"email"`
	Password             string                `json:"password"`
	Content              string                `json:"name"`
	Slug                 string                `json:"slug"`
	PublishedOn          time.Time             `json:"publishedOn"`
	Description          string                `json:"description"`
	Categories           []Category            `bun:"m2m:news_to_categories" json:"categories"`
	Tags                 []Tag                 `bun:"m2m:news_to_tags" json:"tags"`
	Likes                []*NewsLike           `bun:"rel:has-many" json:"likes"`
	PreviewThumbnailFile *PreviewThumbnailFile `bun:"rel:has-one" json:"previewThumbnailFile"`
}
