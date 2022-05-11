package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Chat struct {
	bun.BaseModel `bun:"chats,alias:chats"`
	ID            uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	User          *User         `bun:"rel:belongs-to" json:"user"`
	UserID        uuid.NullUUID `bun:"type:uuid" json:"userId"`
}

type Chats []*Chat
