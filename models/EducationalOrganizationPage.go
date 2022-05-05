package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type EducationalOrganizationPage struct {
	bun.BaseModel `bun:"pages,alias:pages"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Page          *Page     `bun:"rel:belongs-to" json:"page"`
	PageID        uuid.UUID `bun:"type:uuid,nullzero,default:NULL" json:"pageID"`
}

type EducationalOrganizationPages []*EducationalOrganizationPage

func (i EducationalOrganizationPages) SetForeignKeys() {
	for index := range i {
		i[index].PageID = i[index].Page.ID
	}
}

func (i EducationalOrganizationPages) GetPages() Pages {
	items := make(Pages, 0)
	for _, item := range i {
		items = append(items, item.Page)
	}
	return items
}
