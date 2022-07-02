package models

import (
	"github.com/google/uuid"
)

type DivisionComment struct {
	ID         uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	DivisionID uuid.NullUUID `bun:"type:uuid" json:"divisionId"`
	CommentID  uuid.UUID     `bun:"type:uuid" json:"commentId"`
	Division   *Division     `bun:"rel:belongs-to" json:"division"`
	Comment    *Comment      `bun:"rel:belongs-to" json:"comment"`
}

type DivisionComments []*DivisionComment

func (item *DivisionComment) SetForeignKeys() {
	item.CommentID = item.Comment.ID
}
