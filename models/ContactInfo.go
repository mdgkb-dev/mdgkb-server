package models

import (
	"github.com/google/uuid"
)

type ContactInfo struct {
	ID                        uuid.UUID        `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	EmailsForDelete           []uuid.UUID      `bun:"-" json:"emailsForDelete"`
	Emails                    Emails           `bun:"rel:has-many" json:"emails"`
	PostAddressesForDelete    []uuid.UUID      `bun:"-" bun:"rel:has-many" json:"postAddressesForDelete"`
	PostAddresses             PostAddresses    `bun:"rel:has-many" json:"postAddresses"`
	TelephoneNumbersForDelete []uuid.UUID      `bun:"-" json:"telephoneNumbersForDelete"`
	TelephoneNumbers          TelephoneNumbers `bun:"rel:has-many" json:"telephoneNumbers"`
	WebsitesForDelete         []uuid.UUID      `bun:"-" json:"websitesForDelete"`
	Websites                  Websites         `bun:"rel:has-many" json:"websites"`
}

type ContactInfos []*ContactInfo

func (item *ContactInfo) SetIdForChildren() {
	for i := range item.Emails {
		item.Emails[i].ContactInfoId = item.ID
	}
	for i := range item.PostAddresses {
		item.PostAddresses[i].ContactInfoId = item.ID
	}
	for i := range item.TelephoneNumbers {
		item.TelephoneNumbers[i].ContactInfoId = item.ID
	}
	for i := range item.Websites {
		item.Websites[i].ContactInfoId = item.ID
	}
}

func (items ContactInfos) SetIdForChildren() {
	for i := range items {
		items[i].SetIdForChildren()
	}
}

func (items ContactInfos) GetEmails() Emails {
	itemsForGet := make(Emails, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].Emails...)
	}
	return itemsForGet
}

func (items ContactInfos) GetPostAddresses() PostAddresses {
	itemsForGet := make(PostAddresses, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].PostAddresses...)
	}
	return itemsForGet
}

func (items ContactInfos) GetTelephoneNumbers() TelephoneNumbers {
	itemsForGet := make(TelephoneNumbers, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].TelephoneNumbers...)
	}
	return itemsForGet
}

func (items ContactInfos) GetWebsites() Websites {
	itemsForGet := make(Websites, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].Websites...)
	}
	return itemsForGet
}
