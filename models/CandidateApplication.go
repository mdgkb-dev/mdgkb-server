package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type CandidateApplication struct {
	bun.BaseModel `bun:"candidate_applications,alias:candidate_applications"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `

	CandidateExam   *CandidateExam `bun:"rel:belongs-to" json:"candidateExam"`
	CandidateExamID uuid.NullUUID  `bun:"type:uuid,nullzero,default:NULL" json:"candidateExamId"`

	CandidateApplicationSpecializations          CandidateApplicationSpecializations `bun:"rel:has-many" json:"candidateApplicationSpecializations"`
	CandidateApplicationSpecializationsForDelete []uuid.UUID                         `bun:"-" json:"candidateApplicationSpecializationsForDelete"`

	FormValue   *FormValue    `bun:"rel:belongs-to" json:"formValue"`
	FormValueID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"formValueId"`
}

type CandidateApplications []*CandidateApplication

func (item *CandidateApplication) SetForeignKeys() {
	item.CandidateExamID = item.CandidateExam.ID
	item.FormValueID = item.FormValue.ID
}

func (item *CandidateApplication) SetFilePath(fileID *string) *string {
	path := item.FormValue.SetFilePath(fileID)
	return path
}

func (item *CandidateApplication) SetIdForChildren() {
	for i := range item.CandidateApplicationSpecializations {
		item.CandidateApplicationSpecializations[i].CandidateApplicationID = item.ID
	}
}
