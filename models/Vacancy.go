package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Vacancy struct {
	bun.BaseModel    `bun:"vacancies,alias:vacancies"`
	ID               uuid.UUID        `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Title            string           `json:"title"`
	Description      string           `json:"description"`
	Specialization   string           `json:"specialization"`
	Salary           string           `json:"salary"`
	Archived         bool             `json:"archived"`
	Requirements     string           `json:"requirements"`
	Experience       string           `json:"experience"`
	Duties           string           `json:"duties"`
	Schedule         string           `json:"schedule"`
	Date             time.Time        `bun:"vacancy_date" json:"date"`
	VacancyResponses VacancyResponses `bun:"rel:has-many" json:"vacancyResponses"`

	Division   *Division     `bun:"rel:belongs-to" json:"division"`
	DivisionId uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"divisionId,omitempty"`
}

type Vacancies []*Vacancy
