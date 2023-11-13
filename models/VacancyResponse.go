package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type VacancyResponse struct {
	bun.BaseModel `bun:"vacancy_responses,select:vacancy_responses_view,alias:vacancy_responses_view"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`

	Vacancy   *Vacancy  `bun:"rel:belongs-to" json:"vacancy"`
	VacancyID uuid.UUID `bun:"type:uuid"  json:"vacancyId"`

	FormValue   *FormValue    `bun:"rel:belongs-to" json:"formValue"`
	FormValueID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"formValueId"`

	User   *User         `bun:"rel:belongs-to" json:"user"`
	UserID uuid.NullUUID `bun:"user_id,nullzero,type:uuid" json:"userId"`

	FormStatusID string `bun:"-" json:"formStatusId"`
	FullName     string `bun:"-" json:"fullName"`
	Title        string `bun:"-" json:"title"`
	Email        string `bun:"-" json:"email"`
	CreatedAt    string `bun:"-" json:"createdAt"`
}

type VacancyResponses []*VacancyResponse

type VacancyResponsesWithCount struct {
	VacancyResponses VacancyResponses `json:"items"`
	Count            int              `json:"count"`
}

func (item *VacancyResponse) SetForeignKeys() {
	item.FormValueID = item.FormValue.ID
}

func (item *VacancyResponse) SetIDForChildren() {
}

func (item *VacancyResponse) SetFilePath(fileID *string) *string {
	path := item.FormValue.SetFilePath(fileID)
	return path
}
