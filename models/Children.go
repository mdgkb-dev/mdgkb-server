package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Child struct {
	bun.BaseModel `bun:"children,alias:children"`
	ID            uuid.UUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Human         *Human        `bun:"rel:belongs-to" json:"human"`
	HumanID       uuid.UUID `bun:"type:uuid" json:"humanId"`

	User         *User        `bun:"rel:belongs-to" json:"user"`
	UserID       uuid.UUID `bun:"type:uuid" json:"userId"`
}

type Children []*Child

func (i *Child) SetForeignKeys() {
	i.HumanID = i.Human.ID
}

func (items Children) SetForeignKeys()  {
	for i := range items {
		items[i].SetForeignKeys()
	}
}

func (items Children) GetHumans() Humans {
	itemsForGet := make(Humans, len(items))
	for i := range items {
		itemsForGet[i] = items[i].Human
	}
	return itemsForGet
}