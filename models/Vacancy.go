package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Vacancy struct {
	bun.BaseModel       `bun:"vacancies,alias:vacancies"`
	ID                  uuid.UUID           `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Title               string              `json:"title"`
	Slug                string              `json:"slug"`
	Specialization      string              `json:"specialization"`
	MinSalary           int                 `json:"minSalary"`
	MaxSalary           int                 `json:"maxSalary"`
	SalaryComment       string              `json:"salaryComment"`
	Archived            bool                `json:"archived"`
	Experience          string              `json:"experience"`
	Schedule            string              `json:"schedule"`
	Date                time.Time           `bun:"vacancy_date" json:"date"`
	VacancyResponses    VacancyResponses    `bun:"rel:has-many" json:"vacancyResponses"`
	VacancyDuties       VacancyDuties       `bun:"rel:has-many" json:"vacancyDuties"`
	VacancyRequirements VacancyRequirements `bun:"rel:has-many" json:"vacancyRequirements"`

	Division   *Division     `bun:"rel:belongs-to" json:"division"`
	DivisionId uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"divisionId,omitempty"`

	ContactInfo   *ContactInfo `bun:"rel:belongs-to" json:"contactInfo"`
	ContactInfoID uuid.UUID    `bun:"type:uuid" json:"contactInfoId"`

	ContactDoctor   *Doctor   `bun:"rel:belongs-to" json:"contactDoctor"`
	ContactDoctorID uuid.UUID `bun:"type:uuid" json:"contactDoctorId"`
}

type Vacancies []*Vacancy
