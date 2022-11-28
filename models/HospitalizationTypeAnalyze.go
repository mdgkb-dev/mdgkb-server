package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type HospitalizationTypeAnalyze struct {
	bun.BaseModel         `bun:"hospitalization_type_analyzes,alias:hospitalization_type_analyzes"`
	ID                    uuid.UUID            `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name                  string               `json:"name"`
	Children              bool                 `json:"children"`
	DurationDays          uint8                `json:"durationDays"`
	Order                 uint8                `bun:"item_order" json:"order"`
	HospitalizationType   *HospitalizationType `bun:"rel:belongs-to" json:"hospitalizationType"`
	HospitalizationTypeID uuid.NullUUID        `bun:"type:uuid" json:"hospitalizationTypeId"`
}

type HospitalizationTypeAnalyzes []*HospitalizationTypeAnalyze
