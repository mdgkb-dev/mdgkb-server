package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	bun.BaseModel `bun:"users,alias:users"`
	ID            uuid.UUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Email         string        `json:"email"`
	Password      string        `json:"password"`
	Human         *Human        `bun:"rel:belongs-to" json:"human"`
	HumanID       uuid.UUID `bun:"type:uuid" json:"humanId"`
	Questions          Questions `bun:"rel:has-many" json:"questions"`

	Children Children `bun:"rel:has-many" json:"children"`
	ChildrenForDelete []uuid.UUID `bun:"-" json:"childrenForDelete"`
}

type Users []*User

func (i *User) GenerateHashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(i.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	pass := string(hash)
	i.Password = pass
	return nil
}

func (i *User) CompareWithHashPassword(password *string) bool {
	p, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}
	return bcrypt.CompareHashAndPassword(p, []byte(*password)) == nil
}

func (i *User) SetForeignKeys() {
	i.HumanID = i.Human.ID
}

func (i *User) SetIdForChildren() {
	for index := range i.Children {
		i.Children[index].UserID = i.ID
	}
}