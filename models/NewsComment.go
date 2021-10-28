package models

import (
	"github.com/google/uuid"
)

type NewsComment struct {
	ID        uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	NewsID    uuid.UUID `bun:"type:uuid" json:"newsId"`
	CommentID uuid.UUID `bun:"type:uuid" json:"commentId"`
	News      *News     `bun:"rel:belongs-to" json:"news"`
	Comment   *Comment  `bun:"rel:belongs-to" json:"comment"`
}

type NewsComments []*NewsComment
