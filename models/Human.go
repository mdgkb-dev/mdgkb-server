package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helpers/uploader"
	baseModels "github.com/pro-assistance/pro-assister/models"
	"github.com/uptrace/bun"
	// "time"
)

type Human struct {
	bun.BaseModel `bun:"humans,alias:humans"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name" validate:"required,min=1,max=100"`
	Surname       string        `json:"surname" validate:"required,min=1,max=100"`
	Patronymic    string        `json:"patronymic" validate:"required,min=1,max=100"`
	Snils         string        `json:"snils"`

	PostIndex string `json:"postIndex"`
	Address   string `json:"address"`

	IsMale      bool       `json:"isMale"`
	DateBirth   *time.Time `json:"dateBirth"`
	PlaceBirth  string     `json:"placeBirth"`
	Citizenship string     `json:"citizenship"`
	Slug        string     `json:"slug"`

	CarNumber string `json:"carNumber"`
	CarModel  string `json:"carModel"`

	Photo   *FileInfo     `bun:"rel:belongs-to" json:"photo"`
	PhotoID uuid.NullUUID `bun:"type:uuid" json:"photoId"`

	PhotoMini   *FileInfo2    `bun:"rel:belongs-to" json:"photoMini"`
	PhotoMiniID uuid.NullUUID `bun:"type:uuid" json:"photoMiniId"`

	Contact   *baseModels.Contact `bun:"rel:belongs-to" json:"contact"`
	ContactID uuid.NullUUID       `bun:"type:uuid" json:"contactId"`
}

type Humans []*Human

func (item *Human) SetForeignKeys() {
	if item.Contact != nil {
		item.ContactID = item.Contact.ID
	}
	if item.Photo != nil {
		item.PhotoID = item.Photo.ID
	}
	// if item.PhotoMini != nil {
	// 	item.PhotoMiniID = item.PhotoMini.ID
	// }
}

func (items Humans) SetForeignKeys() {
	for i := range items {
		items[i].SetForeignKeys()
	}
}

func (item *Human) GetFullName() string {
	return fmt.Sprintf("%s %s %s", item.Surname, item.Name, item.Patronymic)
}

func (items Humans) GetContacts() baseModels.Contacts {
	itemsForGet := make(baseModels.Contacts, len(items))
	for i := range items {
		if items[i].Contact != nil {
			itemsForGet[i] = items[i].Contact
		}
	}
	return itemsForGet
}

func (items Humans) GetPhotos() FileInfos {
	itemsForGet := make(FileInfos, len(items))
	for i := range items {
		if items[i].Contact != nil {
			itemsForGet[i] = items[i].Photo
		}
	}
	return itemsForGet
}

func (items Humans) GetFileInfos() FileInfos {
	itemsForGet := make(FileInfos, 0)

	for _, item := range items {
		itemsForGet = append(itemsForGet, item.Photo)
		// p := FileInfo(*item.PhotoMini)
		// itemsForGet = append(itemsForGet, &p)
	}
	return itemsForGet
}

func (item *Human) SetFilePath(fileID *string) *string {
	if item.Photo.ID.UUID.String() == *fileID {
		item.Photo.FileSystemPath = uploader.BuildPath(fileID)
		return &item.Photo.FileSystemPath
	}
	// if item.PhotoMini.ID.UUID.String() == *fileID {
	// 	item.PhotoMini.FileSystemPath = uploader.BuildPath(fileID)
	// 	return &item.PhotoMini.FileSystemPath
	// }
	return nil
}

func (item *Human) GetFullAddress() string {
	// return item.Contact.Address.GetFullAddress()
	return item.Address
}
