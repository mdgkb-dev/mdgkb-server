package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Question struct {
	bun.BaseModel    `bun:"questions,alias:questions"`
	ID               uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Question         string    `json:"question"`
	OriginalQuestion string    `json:"originalQuestion"`
	Answer           string    `json:"answer"`
	OriginalAnswer   string    `json:"originalAnswer"`
	Published        bool      `json:"published"`
	Date             time.Time `bun:"question_date" json:"date"`
	User             *User     `bun:"rel:belongs-to" json:"user"`
	UserID           uuid.UUID `bun:"type:uuid" json:"userId"`
	IsNew            bool      `json:"isNew"`
	AnswerIsRead     bool      `json:"answerIsRead"`
}

type Questions []*Question
