package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Postgraduate struct {
	bun.BaseModel    `bun:"postgraduate,alias:postgraduate"`
	ID               uuid.UUID       `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	SpecializationID uuid.UUID       `bun:"type:uuid" json:"specializationId"`
	Specialization   *Specialization `bun:"rel:belongs-to" json:"specialization"`
}

type Postgraduates []*Postgraduate
