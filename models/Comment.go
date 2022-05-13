package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID               uuid.UUID        `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	UserId           uuid.UUID        `bun:"type:uuid" json:"userId"`
	Text             string           `json:"text"`
	ModChecked       bool             `json:"modChecked"`
	Positive         bool             `json:"positive"`
	Answer           string           `json:"answer"`
	PublishedOn      time.Time        `bun:"default:current_timestamp" json:"publishedOn"`
	User             *User            `bun:"rel:belongs-to" json:"user"`
	Rating           float32          `json:"rating"`
	NewsComment      *NewsComment     `bun:"rel:has-one" json:"newsComment"`
	DoctorComments   DoctorComments   `bun:"rel:has-many" json:"doctorComment"`
	DivisionComments DivisionComments `bun:"rel:has-many" json:"divisionComment"`
	PageComments     PageComments     `bun:"rel:has-many" json:"pageComment"`
}

type Comments []*Comment
