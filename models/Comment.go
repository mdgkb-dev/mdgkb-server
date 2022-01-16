package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID              uuid.UUID        `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	UserId          uuid.UUID        `bun:"type:uuid" json:"userId"`
	Text            string           `json:"text"`
	ModChecked      bool             `json:"modChecked"`
	Positive        bool             `json:"positive"`
	PublishedOn     time.Time        `bun:"default:current_timestamp" json:"publishedOn"`
	User            *User            `bun:"rel:belongs-to" json:"user"`
	Rating          float32          `json:"rating"`
	NewsComment     *NewsComment     `bun:"rel:has-one" json:"newsComment"`
	DoctorComments  *DoctorComments  `bun:"rel:has-many" json:"doctorComment"`
	DivisionComment *DivisionComment `bun:"rel:has-one" json:"divisionComment"`
	PageComment     *PageComment     `bun:"rel:has-one" json:"pageComment"`
}

type Comments []*Comment
