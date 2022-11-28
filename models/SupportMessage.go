package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type SupportMessage struct {
	bun.BaseModel `bun:"support_messages,alias:support_messages"`
	ID            uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Theme         string        `json:"theme"`
	Question      string        `json:"question"`
	Answer        string        `json:"answer"`
	Date          string        `bun:"support_message_date" json:"date"`
	User          *User         `bun:"rel:belongs-to" json:"user"`
	UserID        uuid.NullUUID `bun:"type:uuid" json:"userId"`
	IsNew         bool          `json:"isNew"`
}

type SupportMessages []*SupportMessage

func (item *SupportMessage) SetForeignKeys() {
	if item.User != nil {
		item.UserID = item.User.ID
	}
}
