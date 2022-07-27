package models

import (
	"time"

	"github.com/uptrace/bun"

	"github.com/google/uuid"
)

type Comment struct {
	bun.BaseModel `bun:"comments,alias:comments"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	UserID        uuid.UUID `bun:"type:uuid" json:"userId"`
	Text          string    `json:"text"`
	ModChecked    bool      `json:"modChecked"`
	Positive      bool      `json:"positive"`

	Answer          string           `json:"answer"`
	PublishedOn     time.Time        `bun:"default:current_timestamp" json:"publishedOn"`
	User            *User            `bun:"rel:belongs-to" json:"user"`
	Rating          float32          `json:"rating"`
	NewsComment     *NewsComment     `bun:"rel:has-one" json:"newsComment"`
	DoctorComment   *DoctorComment   `bun:"rel:has-one" json:"doctorComment"`
	DivisionComment *DivisionComment `bun:"rel:has-one" json:"divisionComment"`
	PageComment     *PageComment     `bun:"rel:has-one" json:"pageComment"`
}

type Comments []*Comment
