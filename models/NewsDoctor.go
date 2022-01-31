package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type NewsDoctor struct {
	bun.BaseModel `bun:"news_doctors,alias:news_doctors"`
	ID            uuid.UUID     `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	DoctorID      uuid.NullUUID `bun:"type:uuid" json:"doctorId"`
	Doctor        *Doctor       `bun:"rel:belongs-to" json:"doctor"`
	NewsID        uuid.UUID     `bun:"type:uuid" json:"newsId"`
	News          *News         `bun:"rel:belongs-to" json:"news"`
}

type NewsDoctors []*NewsDoctor
