package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type NewsDivision struct {
	bun.BaseModel `bun:"news_divisions,alias:news_divisions"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	DivisionID    uuid.NullUUID `bun:"type:uuid" json:"divisionId"`
	Division      *Division     `bun:"rel:belongs-to" json:"division"`
	NewsID        uuid.NullUUID `bun:"type:uuid" json:"newsId"`
	News          *News         `bun:"rel:belongs-to" json:"news"`
}

type NewsDivisions []*NewsDivision
