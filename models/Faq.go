package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Faq struct {
	bun.BaseModel `bun:"faqs,alias:faqs"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Order         uint      `bun:"faq_order" json:"order"`
	Question      string    `json:"question"`
	Answer        string    `json:"answer"`
}

type Faqs []*Faq

type FaqsWithDelete struct {
	Faqs          Faqs        `json:"faqs"`
	FaqsForDelete []uuid.UUID `json:"faqsForDelete"`
}
