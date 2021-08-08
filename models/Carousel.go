package models

import (
	"github.com/google/uuid"
)

type Carousel struct {
	ID                      uuid.UUID        `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Title                   string           `json:"title"`
	SystemKey               string           `json:"systemKey"`
	CarouselSlides          []*CarouselSlide `bun:"rel:has-many" json:"carouselSlides"`
	CarouselSlidesNames     []string         `bun:"-" json:"carouselSlidesNames"`
	CarouselSlidesForDelete []string         `bun:"-" json:"carouselSlidesForDelete"`
}

type CarouselSlide struct {
	ID          uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	Title       string    `json:"title"`
	CarouselID  uuid.UUID `bun:"type:uuid" json:"carouselId"`
	Content     string    `json:"content"`
	Link        string    `json:"link"`
	ButtonShow  bool      `json:"buttonShow"`
	ButtonColor string    `json:"buttonColor"`
	FileInfo    *FileInfo `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoId  uuid.UUID `bun:"type:uuid" json:"fileInfoId"`
}
