package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Chat struct {
	bun.BaseModel         `bun:"chats,alias:chats"`
	ID                    uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	ChatMessages          ChatMessages  `bun:"rel:has-many" json:"chatMessages"`
	ChatMessagesForDelete []uuid.UUID   `bun:"-" json:"chatMessagesForDelete"`
}

type Chats []*Chat
type ChatsWithCount struct {
	Chats Chats `json:"items"`
	Count int   `json:"count"`
}
