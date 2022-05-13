package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type EducationYear struct {
	bun.BaseModel `bun:"education_years,alias:education_years"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Year          time.Time     `json:"year"`
	Active        bool          `json:"active"`
}

type EducationYears []*EducationYear
