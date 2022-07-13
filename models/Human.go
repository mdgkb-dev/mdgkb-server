package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"github.com/uptrace/bun"
	"time"
	// "time"
)

type Human struct {
	bun.BaseModel `bun:"humans,alias:humans"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Surname       string        `json:"surname"`
	Patronymic    string        `json:"patronymic"`
	Snils         string        `json:"snils"`

	PostIndex string `json:"postIndex"`
	Address   string `json:"address"`

	IsMale      bool       `json:"isMale"`
	DateBirth   *time.Time `json:"dateBirth"`
	PlaceBirth  string     `json:"placeBirth"`
	Citizenship string     `json:"citizenship"`
	Slug        string     `json:"slug"`
	
	CarNumber   string     `json:"carNumber"`
	CarModel    string     `json:"carModel"`

	Photo   *FileInfo     `bun:"rel:belongs-to" json:"photo"`
	PhotoID uuid.NullUUID `bun:"type:uuid" json:"photoId"`

	ContactInfo   *ContactInfo `bun:"rel:belongs-to" json:"contactInfo"`
	ContactInfoID uuid.UUID    `bun:"type:uuid" json:"contactInfoId"`
}

type Humans []*Human

func (item *Human) SetForeignKeys() {
	item.ContactInfoID = item.ContactInfo.ID
	item.PhotoID = item.Photo.ID
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
		if items[i].ContactInfo != nil {
			itemsForGet[i] = items[i].ContactInfo
		}
	}
	return itemsForGet
}

func (items Humans) GetPhotos() FileInfos {
	itemsForGet := make(FileInfos, len(items))
	for i := range items {
		if items[i].ContactInfo != nil {
			itemsForGet[i] = items[i].Photo
		}
	}
	return itemsForGet
}

func (item *Human) SetFilePath(fileID *string) *string {
	if item.Photo.ID.UUID.String() == *fileID {
		item.Photo.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.Photo.FileSystemPath
	}
	return nil
}
