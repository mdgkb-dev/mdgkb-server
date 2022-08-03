package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"github.com/uptrace/bun"
)

type Question struct {
	bun.BaseModel    `bun:"questions,alias:questions"`
	ID               uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Theme            string        `json:"theme"`
	Question         string        `json:"question"`
	OriginalQuestion string        `json:"originalQuestion"`
	Answer           string        `json:"answer"`
	OriginalAnswer   string        `json:"originalAnswer"`
	PublishAgreement bool          `json:"publishAgreement"`
	Published        bool          `json:"published"`
	Answered         bool          `json:"answered"`
	Date             time.Time     `bun:"question_date" json:"date"`
	User             *User         `bun:"rel:belongs-to" json:"user"`
	UserID           uuid.NullUUID `json:"type:uuid" json:"userId"`
	IsNew            bool          `json:"isNew"`
	AnswerIsRead     bool          `json:"answerIsRead"`

	File   *FileInfo     `bun:"rel:belongs-to" json:"file"`
	FileID uuid.NullUUID `json:"fileId"`
}

type Questions []*Question

func (item *Question) SetForeignKeys() {
	if item.User != nil {
		item.UserID = item.User.ID
	}
	if item.File != nil {
		item.FileID = item.File.ID
	}
}

// func (item *Question) SetFilePath(fileID *string) *string {
// 	path := item....SetFilePath(fileID);
// 	return path
// }

func (item *Question) SetFilePath(fileID *string) *string {
	if item.File.ID.UUID.String() == *fileID {
		item.File.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.File.FileSystemPath
	}
	return nil
}
