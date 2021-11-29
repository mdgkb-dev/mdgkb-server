package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"

	// "time"
)

type Human struct {
	bun.BaseModel `bun:"humans,alias:humans"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`
	Surname       string    `json:"surname"`
	Patronymic    string    `json:"patronymic"`
	IsMale        bool      `json:"isMale"`
	DateBirth     *time.Time `json:"dateBirth,omitempty"`
	Slug string `json:"slug"`

	ContactInfo   *ContactInfo `bun:"rel:belongs-to" json:"contactInfo"`
	ContactInfoID uuid.UUID    `bun:"type:uuid" json:"contactInfoId"`
}

type Humans []*Human

func (item *Human) SetForeignKeys() {
	item.ContactInfoID = item.ContactInfo.ID
}

func (items Humans) SetForeignKeys() {
	for i := range items {
		items[i].SetForeignKeys()
	}
}

func (item *Human) GetFullName() string {
	return fmt.Sprintf("%s %s %s", item.Surname, item.Name, item.Patronymic)
}



func (items Humans) GetContactInfos() ContactInfos {
	itemsForGet := make(ContactInfos, len(items))
	for i := range items {
		itemsForGet[i] = items[i].ContactInfo
	}
	return itemsForGet
}