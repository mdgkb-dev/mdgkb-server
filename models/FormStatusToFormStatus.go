package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type FormStatusToFormStatus struct {
	bun.BaseModel     `bun:"form_status_to_form_statuses,alias:form_status_to_form_statuses"`
	ID                uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	FormStatus        *FormStatus   `bun:"rel:belongs-to" json:"formStatus"`
	FormStatusID      uuid.NullUUID `bun:"type:uuid" json:"formStatusId"`
	ChildFormStatus   *FormStatus   `bun:"rel:belongs-to" json:"childFormStatus"`
	ChildFormStatusID uuid.NullUUID `bun:"type:uuid" json:"childFormStatusId"`
}

type FormStatusToFormStatuses []*FormStatusToFormStatus
