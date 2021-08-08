package models

import (
	"github.com/google/uuid"
	"time"
)

type News struct {
	ID                  uuid.UUID      `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Status              string         `json:"status"`
	Title               string         `json:"title"`
	PreviewText         string         `json:"preview_text"`
	Content             string         `json:"content"`
	Slug                string         `json:"slug"`
	PublishedOn         time.Time      `json:"publishedOn"`
	Description         string         `json:"description"`
	Categories          []Category     `bun:"m2m:news_to_categories" json:"categories"`
	Tags                []Tag          `bun:"m2m:news_to_tags" json:"tags"`
	NewsLikes           []*NewsLike    `bun:"rel:has-many" json:"newsLikes"`
	NewsComments        []*NewsComment `bun:"rel:has-many" json:"newsComments"`
	NewsImages          []*NewsImage   `bun:"rel:has-many" json:"newsImages"`
	NewsImagesForDelete []string       `bun:"-" json:"newsImagesForDelete"`
	NewsImagesNames     []string       `bun:"-" json:"newsImagesNames"`
	FileInfo            *FileInfo      `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoId          uuid.UUID      `bun:"type:uuid" json:"fileInfoId"`
	MainImage           *FileInfo      `bun:"rel:belongs-to" json:"mainImage"`
	MainImageID         uuid.UUID      `bun:"type:uuid" json:"mainImageId"`
	NewsViews           []*NewsViews   `bun:"rel:has-many" json:"newsViews"`
	ViewsCount          int            `bun:"-" json:"viewsCount"`
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
	NewsId uuid.UUID `bun:"type:uuid" json:"newsId" `
	News   *News     `bun:"rel:has-one" json:"news" `
	TagId  uuid.UUID `bun:"type:uuid" json:"tagId" `
	Tag    *Tag      `bun:"rel:has-one" json:"tag" `
}
