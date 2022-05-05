package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type NewsComment struct {
	bun.BaseModel `bun:"news_comments,alias:news_comments"`
	ID            uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	NewsID        uuid.NullUUID `bun:"type:uuid" json:"newsId"`
	CommentID     uuid.UUID     `bun:"type:uuid" json:"commentId"`
	News          *News         `bun:"rel:belongs-to" json:"news"`
	Comment       *Comment      `bun:"rel:belongs-to" json:"comment"`
}

type NewsComments []*NewsComment
