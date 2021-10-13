package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type VacancyResponse struct {
	bun.BaseModel `bun:"vacancy_responses,alias:vacancy_responses"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	ResponseDate  time.Time `json:"responseDate"`
	CoverLetter   string    `json:"coverLetter"`
	Vacancy       *Vacancy  `bun:"rel:belongs-to" json:"vacancy"`
	VacancyID     uuid.UUID `bun:"type:uuid"  json:"vacancyId"`
	Viewed        bool      `json:"viewed"`

	Human         *Human       `bun:"rel:belongs-to" json:"human"`
	HumanID       uuid.UUID    `bun:"type:uuid" json:"humanID"`
	ContactInfo   *ContactInfo `bun:"rel:belongs-to" json:"contactInfo"`
	ContactInfoId uuid.UUID    `bun:"type:uuid" json:"contactInfoId"`
}

type VacancyResponses []*VacancyResponse

func (item *VacancyResponse) SetForeignKeys() {
	item.ContactInfoId = item.ContactInfo.ID
}
