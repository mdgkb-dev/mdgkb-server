package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Email    string    `json:"email"`
	Password string    `json:"password"`
}
