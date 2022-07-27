package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Department struct {
	bun.BaseModel `bun:"departments,select:departments,alias:departments"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Head          *Head         `bun:"rel:belongs-to" json:"head"`
	HeadID        uuid.NullUUID `bun:"type:uuid" json:"headId"`
	IsDivision    bool          `json:"isDivision"`

	Division   *Division     `bun:"rel:belongs-to" json:"division"`
	DivisionID uuid.NullUUID `bun:"type:uuid" json:"divisionId,omitempty"`

	Name string `json:"name"`
}

type Departments []*Department
