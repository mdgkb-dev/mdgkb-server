package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Diploma struct {
	bun.BaseModel     `bun:"diplomas,alias:diplomas"`
	ID                uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	UniversityName    string        `json:"universityName"`
	UniversityEndDate *time.Time    `json:"universityEndDate"`
	Number            string        `json:"number"`
	Series            string        `json:"series"`
	Date              *time.Time    `json:"date"`
	SpecialityCode    string        `json:"specialityCode"`
	Speciality        string        `json:"speciality"`
}
