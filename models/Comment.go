package models

import (
	"time"

	"github.com/uptrace/bun"

	"github.com/google/uuid"
)

type Comment struct {
	bun.BaseModel `bun:"comments,alias:comments"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	UserID        uuid.NullUUID `bun:"type:uuid" json:"userId"`

	ItemID uuid.NullUUID `bun:"type:uuid" json:"itemId"`
	Domen  string        `json:"domen"`

	Text       string `json:"text"`
	ModChecked bool   `json:"modChecked"`
	Positive   bool   `json:"positive"`

	Answer      string    `json:"answer"`
	PublishedOn time.Time `bun:"default:current_timestamp" json:"publishedOn"`
	User        *User     `bun:"rel:belongs-to" json:"user"`
	Rating      float32   `json:"rating"`
}

type Comments []*Comment

type CommentsWithCount struct {
	Comments Comments `json:"items"`
	Count    int      `json:"count"`
}
