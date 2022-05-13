package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type NewsSlideButton struct {
	bun.BaseModel   `bun:"news_slide_buttons,alias:news_slide_buttons"`
	ID              uuid.UUID  `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name            string     `json:"name"`
	BackgroundColor string     `json:"backgroundColor"`
	Color           string     `json:"color"`
	Link            string     `json:"link"`
	Order           uint       `bun:"news_slide_button_order" json:"order"`
	NewsSlide       *NewsSlide `bun:"rel:belongs-to" json:"newsSlide"`
	NewsSlideId     uuid.UUID  `bun:"type:uuid" json:"newsSlideId"`
	BorderColor     string     `json:"borderColor"`
	Shadow          bool       `json:"shadow"`
}

type NewsSlideButtons []*NewsSlideButton
