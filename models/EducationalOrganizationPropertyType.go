package models

import (
	"github.com/google/uuid"
)

type EducationalOrganizationPropertyType struct {
	ID   uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name string    `json:"name"`
}
