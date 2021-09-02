package models

import (
	"github.com/google/uuid"
)

type DoctorComment struct {
	ID        uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	DoctorID  uuid.UUID `bun:"type:uuid" json:"doctorId"`
	CommentId uuid.UUID `bun:"type:uuid" json:"commentId"`
	Doctor    *Doctor   `bun:"rel:belongs-to" json:"doctor"`
	Comment   *Comment  `bun:"rel:belongs-to" json:"comment"`
}
