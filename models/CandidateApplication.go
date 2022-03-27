package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type CandidateApplication struct {
	bun.BaseModel `bun:"candidate_applications,alias:candidate_applications"`
	ID            uuid.NullUUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	CreatedAt     time.Time     `json:"createdAt"`

	CandidateExam   *CandidateExam `bun:"rel:belongs-to" json:"candidateExam"`
	CandidateExamID uuid.NullUUID  `bun:"type:uuid,nullzero,default:NULL" json:"candidateExamId"`

	CandidateApplicationSpecializations          CandidateApplicationSpecializations `bun:"rel:has-many" json:"candidateApplicationSpecializations"`
	CandidateApplicationSpecializationsForDelete []uuid.UUID                         `bun:"-" json:"candidateApplicationSpecializationsForDelete"`

	User   *User     `bun:"rel:belongs-to" json:"user"`
	UserID uuid.UUID `bun:"type:uuid" json:"userId"`

	FieldValues FieldValues `bun:"rel:has-many" json:"fieldValues"`
	IsNew       bool        `json:"isNew"`
}

type CandidateApplications []*CandidateApplication

func (item *CandidateApplication) SetForeignKeys() {
	item.UserID = item.User.ID
	item.CandidateExamID = item.CandidateExam.ID
}

func (item *CandidateApplication) SetFilePath(fileID *string) *string {
	for i := range item.FieldValues {
		filePath := item.FieldValues[i].SetFilePath(fileID)
		if filePath != nil {
			return filePath
		}
	}
	return nil
}

func (item *CandidateApplication) SetIdForChildren() {
	for i := range item.FieldValues {
		item.FieldValues[i].CandidateApplicationID = item.ID
	}
	for i := range item.CandidateApplicationSpecializations {
		item.CandidateApplicationSpecializations[i].CandidateApplicationID = item.ID
	}
}
