package models

import (
	"github.com/uptrace/bun"
	"time"

	"github.com/google/uuid"
)

type Human struct {
	bun.BaseModel `bun:"humans,alias:humans"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`
	Surname       string    `json:"surname"`
	Patronymic    string    `json:"patronymic"`
	IsMale        bool      `json:"isMale"`
	DateBirth     time.Time `bun:"default:current_timestamp" json:"dateBirth"`
}
