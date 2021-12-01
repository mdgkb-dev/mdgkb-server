package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type ChatMessage struct {
	bun.BaseModel `bun:"chat_messages,alias:chat_messages"`
	ID            uuid.UUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	User          *User         `bun:"rel:belongs-to" json:"user"`
	UserID        uuid.NullUUID `bun:"type:uuid" json:"userId"`
	Message       string        `json:"message"`
	IsAnswer      string        `json:"isAnswer"`
	AnswerUser    *User         `bun:"rel:belongs-to" json:"answerUser"`
	AnswerUserID  uuid.NullUUID `bun:"type:uuid" json:"answerUserId"`
	Date          time.Time     `bun:"chat_message_date" json:"date"`
}

type ChatMessages []*ChatMessage
