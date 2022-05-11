package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DoctorComment struct {
	bun.BaseModel `bun:"doctor_comments,alias:doctor_comments"`
	ID            uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	DoctorID      uuid.NullUUID `bun:"type:uuid" json:"doctorId"`
	Doctor        *Doctor       `bun:"rel:belongs-to" json:"doctor"`
	CommentId     uuid.UUID     `bun:"type:uuid" json:"commentId"`
	Comment       *Comment      `bun:"rel:belongs-to" json:"comment"`
}

type DoctorComments []*DoctorComment
