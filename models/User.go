package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func (u *User) GenerateHashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	pass := string(hash)
	u.Password = pass
	return nil
}

func (u *User) CompareWithHashPassword(password *string) bool {
	p, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}
	return bcrypt.CompareHashAndPassword(p, []byte(*password)) == nil
}
