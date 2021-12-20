package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Project struct {
	bun.BaseModel `bun:"projects,alias:projects"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Title         string        `json:"title"`
	Content       string        `json:"content"`
	Slug          string        `json:"slug"`

	ProjectItems          ProjectItems `bun:"rel:has-many" json:"projectItems"`
	ProjectItemsForDelete []uuid.UUID     `bun:"-" json:"projectItemsForDelete"`
}

type Projects []*Project

func (item *Project) SetIdForChildren() {
	for i := range item.ProjectItems {
		item.ProjectItems[i].ProjectID = item.ID
	}
}