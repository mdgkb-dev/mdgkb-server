package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type DocumentForVacancy struct {
	bun.BaseModel      `bun:"documents_for_vacancies,alias:documents_for_vacancies"`
	ID                 uuid.UUID      `bun:"type:uuid,default:uuid_generate_v4()" json:"id" json:"id,omitempty"`
	Documents          DocumentsTypes `bun:"rel:has-many" json:"documents"`
	DocumentsForDelete []uuid.UUID    `bun:"-" json:"documentsForDelete"`
}

type DocumentsForVacancies []*DocumentForVacancy
