package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type VacancyResponse struct {
	bun.BaseModel               `bun:"vacancy_responses,alias:vacancy_responses"`
	ID                          uuid.UUID                   `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	ResponseDate                time.Time                   `json:"responseDate"`
	CoverLetter                 string                      `json:"coverLetter"`
	Vacancy                     *Vacancy                    `bun:"rel:belongs-to" json:"vacancy"`
	VacancyID                   uuid.UUID                   `bun:"type:uuid"  json:"vacancyId"`
	Viewed                      bool                        `json:"viewed"`
	VacancyResponsesToDocuments VacancyResponsesToDocuments `bun:"rel:has-many" json:"vacancyResponsesToDocuments"`

	User   *User     `bun:"rel:belongs-to" json:"user"`
	UserID uuid.UUID `bun:"type:uuid" json:"userId"`
}

type VacancyResponses []*VacancyResponse

func (item *VacancyResponse) SetForeignKeys() {
	item.UserID = item.User.ID
}

func (item *VacancyResponse) SetIdForChildren() {
	if len(item.VacancyResponsesToDocuments) == 0 {
		return
	}
	for i := range item.VacancyResponsesToDocuments {
		item.VacancyResponsesToDocuments[i].VacancyResponseID = item.ID
	}
}

func (item *VacancyResponse) SetFilePath(fileID *string) *string {
	path := item.VacancyResponsesToDocuments.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}
