package models

import (
	"github.com/google/uuid"
)

type Carousel struct {
	ID             uuid.UUID        `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	CarouselSlides []*CarouselSlide `bun:"rel:has-many" json:"carouselSlides"`
}

type CarouselSlide struct {
	ID         uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	Title      string    `json:"title"`
	CarouselID uuid.UUID `bun:"type:uuid" json:"carouselId"`
	Content    string    `json:"content"`
	FileInfo   *FileInfo `bun:"rel:belongs-to" json:"fileInfo"`
	FileInfoId uuid.UUID `bun:"type:uuid" json:"fileInfoId"`
}
