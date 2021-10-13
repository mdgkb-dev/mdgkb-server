package models

import (
	"github.com/google/uuid"
)

type ContactInfo struct {
	ID                        uuid.UUID        `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	EmailsForDelete           []uuid.UUID      `bun:"-" json:"emailsForDelete"`
	Emails                    Emails           `bun:"rel:has-many" json:"emails"`
	PostAddressesForDelete    []uuid.UUID      `bun:"-" bun:"rel:has-many" json:"postAddressesForDelete"`
	PostAddresses             PostAddresses    `bun:"rel:has-many" json:"postAddresses"`
	TelephoneNumbersForDelete []uuid.UUID      `bun:"-" json:"telephoneNumbersForDelete"`
	TelephoneNumbers          TelephoneNumbers `bun:"rel:has-many" json:"telephoneNumbers"`
	WebsitesForDelete         []uuid.UUID      `bun:"-" json:"websitesForDelete"`
	Websites                  Websites         `bun:"rel:has-many" json:"websites"`
}

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
