package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type VacancyResponse struct {
	bun.BaseModel `bun:"vacancy_responses,alias:vacancy_responses"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`

	ResponseDate string `json:"responseDate"`

	CoverLetter string `json:"coverLetter"`

	VacancyId string `json:"vacancyId"`
}

type VacancyResponses []*VacancyResponse
