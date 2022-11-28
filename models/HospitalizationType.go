package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type HospitalizationType struct {
	bun.BaseModel `bun:"hospitalizations_types,alias:hospitalizations_types"`
	ID            uuid.NullUUID  `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Order         uint8          `bun:"hospitalization_type_order" json:"order"`
	Description   string         `json:"description"`
	PolicyType    *PolicyType    `json:"policyType"`
	TreatmentType *TreatmentType `json:"treatmentType"`
	StayType      *StayType      `json:"stayType"`
	ReferralType  *ReferralType  `json:"referralType"`

	FormPattern   *FormPattern  `bun:"rel:belongs-to" json:"formPattern"`
	FormPatternID uuid.NullUUID `bun:"type:uuid" json:"formPatternId"`

	HospitalizationTypeAnalyzes  HospitalizationTypeAnalyzes  `bun:"rel:has-many" json:"hospitalizationTypeAnalyzes"`
	HospitalizationTypeDocuments HospitalizationTypeDocuments `bun:"rel:has-many" json:"hospitalizationTypeDocuments"`
	HospitalizationTypeStages    HospitalizationTypeStages    `bun:"rel:has-many" json:"hospitalizationTypeStages"`
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
