package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type HospitalizationTypeDocument struct {
	bun.BaseModel `bun:"hospitalization_type_documents,alias:hospitalization_type_documents"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name          string    `json:"name"`
	Children      bool      `json:"children"`

	HospitalizationType   *HospitalizationType `bun:"rel:belongs-to" json:"hospitalizationType"`
	HospitalizationTypeID uuid.NullUUID        `bun:"type:uuid" json:"hospitalizationTypeId"`
}

type HospitalizationTypeDocuments []*HospitalizationTypeDocument
