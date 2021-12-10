package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Department struct {
	bun.BaseModel  `bun:"departments,select:departments,alias:departments"`
	ID             uuid.NullUUID      `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Head          *Head         `bun:"rel:belongs-to" json:"head"`
	HeadID        uuid.NullUUID      `bun:"type:uuid" json:"headId"`

	Name string `json:"name"`
}

type Departments []*Department
