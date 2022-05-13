package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type VacancyDuty struct {
	bun.BaseModel `bun:"vacancy_duties,alias:vacancy_duties"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string    `json:"name"`
	Vacancy       *Vacancy  `bun:"rel:belongs-to" json:"vacancy"`
	VacancyID     uuid.UUID `bun:"type:uuid"  json:"vacancyId"`
	Order         uint      `bun:"vacancy_duty_order" json:"order"`
}

type VacancyDuties []*VacancyDuty
