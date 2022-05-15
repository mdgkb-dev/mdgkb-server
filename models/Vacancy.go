package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Vacancy struct {
	bun.BaseModel                `bun:"vacancies,select:vacancies_view,alias:vacancies_view"`
	ID                           uuid.UUID           `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Title                        string              `json:"title"`
	Slug                         string              `json:"slug"`
	Specialization               string              `json:"specialization"`
	MinSalary                    int                 `json:"minSalary"`
	MaxSalary                    int                 `json:"maxSalary"`
	SalaryComment                string              `json:"salaryComment"`
	Active                       bool                `json:"active"`
	Experience                   string              `json:"experience"`
	Schedule                     string              `json:"schedule"`
	Date                         time.Time           `bun:"vacancy_date" json:"date"`
	ResponsesCount               int                 `json:"responsesCount"`
	NewResponsesCount            int                 `json:"newResponsesCount"`
	VacancyResponses             VacancyResponses    `bun:"rel:has-many" json:"vacancyResponses"`
	VacancyDuties                VacancyDuties       `bun:"rel:has-many" json:"vacancyDuties"`
	VacancyDutiesDelete          []uuid.UUID         `bun:"-" json:"vacancyDutiesDelete"`
	VacancyRequirements          VacancyRequirements `bun:"rel:has-many" json:"vacancyRequirements"`
	VacancyRequirementsForDelete []uuid.UUID         `bun:"-" json:"vacancyRequirementsForDelete"`

	Division   *Division     `bun:"rel:belongs-to" json:"division"`
	DivisionId uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"divisionId,omitempty"`

	ContactInfo   *ContactInfo  `bun:"rel:belongs-to" json:"contactInfo"`
	ContactInfoID uuid.NullUUID `bun:"type:uuid" json:"contactInfoId"`

	ContactDoctor   *Doctor       `bun:"rel:belongs-to" json:"contactDoctor"`
	ContactDoctorID uuid.NullUUID `bun:"type:uuid" json:"contactDoctorId"`
}

type Vacancies []*Vacancy

func (item *Vacancy) SetIdForChildren() {
	for i := range item.VacancyDuties {
		item.VacancyDuties[i].VacancyID = item.ID
	}
	for i := range item.VacancyRequirements {
		item.VacancyRequirements[i].VacancyID = item.ID
	}
}
