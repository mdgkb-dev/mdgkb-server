package models

import (
	"github.com/google/uuid"
)

type NewsView struct {
	ID        uuid.UUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	IPAddress string        `json:"ip_address"`
	NewsID    uuid.NullUUID `bun:"type:uuid" json:"newsId"`
	News      *News         `bun:"rel:belongs-to" json:"news"`
}
type NewsViews []*NewsView
