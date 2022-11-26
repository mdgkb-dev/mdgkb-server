package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type HospitalizationType struct {
	bun.BaseModel `bun:"hospitalizations_types,alias:hospitalizations_types"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Paid          bool      `json:"paid"`
	Order         uint8     `bun:"hospitalization_type_order" json:"order"`

	PolicyType    *PolicyType    `json:"policyType"`
	TreatmentType *TreatmentType `json:"treatmentType"`
	StayType      *StayType      `json:"stayType"`
	ReferralType  *ReferralType  `json:"referralType"`

	FormPattern   *FormPattern  `bun:"rel:belongs-to" json:"formPattern"`
	FormPatternID uuid.NullUUID `bun:"type:uuid" json:"formPatternId"`
}

type HospitalizationsTypes []*HospitalizationType

type PolicyType string

const (
	OMS PolicyType = "ОМС"
	DMS PolicyType = "ДМС"
)

type TreatmentType string

const (
	Conservative TreatmentType = "Консервативное"
	Operative    TreatmentType = "Оперативное"
)

type StayType string

const (
	AllDay   StayType = "Круглосуточное"
	ShortDay StayType = "Кратковременное"
)

type ReferralType string

const (
	Moscow ReferralType = "Из поликлинники"
	Region ReferralType = "Из региона"
)
