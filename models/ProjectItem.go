package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type ProjectItem struct {
	bun.BaseModel `bun:"project_items,alias:project_items"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Title         string        `json:"title"`
	Content       string        `json:"content"`

	Project   *Project      `bun:"rel:belongs-to" json:"project"`
	ProjectID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"projectId"`
}

type ProjectItems []*ProjectItem
