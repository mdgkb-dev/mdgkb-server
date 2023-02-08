package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Hospitalization struct {
	bun.BaseModel `bun:"hospitalizations,select:hospitalizations_view,alias:hospitalizations_view"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Date          time.Time `bun:"hospitalization_date" json:"date"`

	FormValue   *FormValue    `bun:"rel:belongs-to" json:"formValue"`
	FormValueID uuid.NullUUID `bun:"type:uuid" json:"formValueId"`

	Division   *Division     `bun:"rel:belongs-to" json:"division"`
	DivisionID uuid.NullUUID `bun:"type:uuid" json:"divisionId"`

	HospitalizationType   *HospitalizationType `bun:"rel:belongs-to" json:"hospitalizationType"`
	HospitalizationTypeID uuid.NullUUID        `bun:"type:uuid" json:"hospitalizationTypeId"`

	HospitalizationsToDocumentTypes HospitalizationsToDocumentTypes `bun:"rel:has-many" json:"hospitalizationsToDocumentTypes"`

	CreatedAt string `bun:"-" json:"createdAt"`
}

type Hospitalizations []*Hospitalization

func (item *Hospitalization) SetForeignKeys() {
	if item.FormValue != nil {
		item.FormValueID = item.FormValue.ID
	}
}

func (item *Hospitalization) SetFilePath(fileID *string) *string {
	path := item.FormValue.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}
