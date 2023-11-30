package models

import (
	"time"

	"github.com/google/uuid"
)

type NewsView struct {
	ID        uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	IPAddress string
	NewsID    uuid.NullUUID `bun:"type:uuid" json:"newsId"`
	News      *News         `bun:"rel:belongs-to" json:"news"`

	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
	Country   string    `json:"country"`
	City      string    `json:"city"`
}

type NewsViews []*NewsView
