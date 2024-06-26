package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Gate struct {
	bun.BaseModel `bun:"gates,alias:gates"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	FormPatternID uuid.NullUUID `bun:"type:uuid" json:"formPatternId"`
	Name          string        `json:"name"`
	Num           uint8         `json:"num"`

	FormPattern        *FormPattern       `bun:"rel:belongs-to" json:"formPattern"`
	VisitsApplications VisitsApplications `bun:"rel:has-many" json:"visitsApplications"`
}

type Gates []*Gate

func (item *Gate) SetForeignKeys() {
	item.FormPatternID = item.FormPattern.ID
}

func (item *Gate) SetFilePath(fileID *string) *string {
	return item.FormPattern.SetFilePath(fileID)
}

func (items Gates) SetForeignKeys() {
	for i := range items {
		items[i].SetForeignKeys()
	}
}

func (items Gates) SetFilePath(fileID *string) *string {
	for i := range items {
		path := items[i].SetFilePath(fileID)
		if path != nil {
			return path
		}
	}
	return nil
}
