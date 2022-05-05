package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type EducationalOrganizationProperty struct {
	bun.BaseModel                         `bun:"educational_organization_properties,alias:educational_organization_properties"`
	ID                                    uuid.UUID                            `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name                                  string                               `json:"name"`
	Value                                 string                               `json:"value"`
	EducationalOrganizationPropertyType   *EducationalOrganizationPropertyType `bun:"rel:belongs-to" json:"educationalOrganizationPropertyType"`
	EducationalOrganizationPropertyTypeID uuid.UUID                            `bun:"type:uuid,nullzero,default:NULL" json:"educationalOrganizationPropertyTypeId,omitempty"`
	Order                                 uint                                 `bun:"educational_organization_property_order" json:"order"`
}

type EducationalOrganizationProperties []*EducationalOrganizationProperty
