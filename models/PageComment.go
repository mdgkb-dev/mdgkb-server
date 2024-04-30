package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PageComment struct {
	bun.BaseModel `bun:"pages_comments,alias:pages_comments"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`

	Page      *Page         `bun:"rel:belongs-to" json:"page"`
	PageID    uuid.NullUUID `bun:"type:uuid" json:"pageId"`
	Comment   *Comment      `bun:"rel:belongs-to" json:"comment"`
	CommentID uuid.NullUUID `bun:"type:uuid" json:"commentId"`
}

type PageComments []*PageComment

func (items PageComments) GetComments() Comments {
	itemsForGet := make(Comments, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.Comment)
	}
	return itemsForGet
}

func (items PageComments) SetForeignKeys() {
	for i := range items {
		items[i].CommentID = items[i].Comment.ID
	}
}
