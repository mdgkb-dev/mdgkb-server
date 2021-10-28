package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	bun.BaseModel `bun:"users,alias:users"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	Human         *Human    `bun:"rel:belongs-to" json:"human"`
	HumanID       uuid.UUID `bun:"type:uuid" json:"humanId"`
}

type Users []*User

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
