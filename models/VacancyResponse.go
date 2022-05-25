package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type VacancyResponse struct {
	bun.BaseModel `bun:"vacancy_responses,alias:vacancy_responses"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`

	Vacancy   *Vacancy  `bun:"rel:belongs-to" json:"vacancy"`
	VacancyID uuid.UUID `bun:"type:uuid"  json:"vacancyId"`

	FormValue   *FormValue    `bun:"rel:belongs-to" json:"formValue"`
	FormValueID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"formValueId"`
}

type VacancyResponses []*VacancyResponse

func (item *VacancyResponse) SetForeignKeys() {
	item.FormValueID = item.FormValue.ID
}

func (item *VacancyResponse) SetIdForChildren() {
}

func (item *VacancyResponse) SetFilePath(fileID *string) *string {
	path := item.FormValue.SetFilePath(fileID)
	return path
}
