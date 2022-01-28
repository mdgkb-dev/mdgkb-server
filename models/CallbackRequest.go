package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type CallbackRequest struct {
	bun.BaseModel `bun:"callback_requests,alias:callback_requests"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Phone         string    `json:"phone"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
}

type CallbackRequests []*CallbackRequest
