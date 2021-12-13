package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Question struct {
	bun.BaseModel    `bun:"questions,alias:questions"`
	ID               uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Theme            string    `json:"theme"`
	Question         string    `json:"question"`
	OriginalQuestion string    `json:"originalQuestion"`
	Answer           string    `json:"answer"`
	OriginalAnswer   string    `json:"originalAnswer"`
	PublishAgreement bool      `json:"publishAgreement"`
	Published        bool      `json:"published"`
	Answered         bool      `json:"answered"`
	Date             time.Time `bun:"question_date" json:"date"`
	User             *User     `bun:"rel:belongs-to" json:"user"`
	UserID           uuid.UUID `bun:"type:uuid" json:"userId"`
	IsNew            bool      `json:"isNew"`
	AnswerIsRead     bool      `json:"answerIsRead"`
}

type Questions []*Question

func (item *Question) SetForeignKeys() {
	if item.User != nil {
		item.UserID = item.User.ID
	}
}
