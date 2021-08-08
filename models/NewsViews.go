package models

import (
	"github.com/google/uuid"
)

type NewsViews struct {
	ID        uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	IPAddress string    `json:"ip_address"`
	NewsID    uuid.UUID `bun:"type:uuid" json:"newsId"`
	News      *News     `bun:"rel:belongs-to" json:"news"`
}
