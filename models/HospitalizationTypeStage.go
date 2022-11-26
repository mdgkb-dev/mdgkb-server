package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type HospitalizationTypeStage struct {
	bun.BaseModel `bun:"hospitalization_type_stages,alias:hospitalization_type_stages"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name          string    `json:"name"`

	HospitalizationType   *HospitalizationType `bun:"rel:belongs-to" json:"hospitalizationType"`
	HospitalizationTypeID uuid.NullUUID        `bun:"type:uuid" json:"hospitalizationTypeId"`
}

type HospitalizationTypeStages []*HospitalizationType
