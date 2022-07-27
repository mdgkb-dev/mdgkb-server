package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DocumentForVacancy struct {
	bun.BaseModel      `bun:"documents_for_vacancies,alias:documents_for_vacancies"`
	ID                 uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id,omitempty"`
	Documents          DocumentTypes `bun:"rel:has-many" json:"documents"`
	DocumentsForDelete []uuid.UUID   `bun:"-" json:"documentsForDelete"`
}

type DocumentsForVacancies []*DocumentForVacancy
