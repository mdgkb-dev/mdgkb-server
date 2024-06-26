package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PaidService struct {
	bun.BaseModel `bun:"paid_services,alias:paid_services"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`
	Price         int       `json:"price"`
	Code          string    `json:"code"`
}

type PaidServices []*PaidService
