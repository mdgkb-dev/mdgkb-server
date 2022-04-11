package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type PublicDocumentType struct {
	bun.BaseModel               `bun:"public_document_types,alias:public_document_types"`
	ID                          uuid.NullUUID                `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	Name                        string                       `json:"name"`
	RouteAnchor                 string                       `json:"routeAnchor"`
	Description                 string                       `json:"description"`
	DocumentTypes               DocumentTypes                `bun:"rel:has-many" json:"documentTypes"`
	DocumentTypesForDelete      []uuid.UUID                  `bun:"-" json:"documentTypesForDelete"`
	EducationPublicDocumentType *EducationPublicDocumentType `bun:"rel:has-one" json:"educationPublicDocumentType"`
}

type PublicDocumentTypes []*PublicDocumentType

func (item PublicDocumentType) SetFilePath(fileID *string) *string {
	for i := range item.DocumentTypes {
		filePath := item.DocumentTypes[i].SetFilePath(fileID)
		if filePath != nil {
			return filePath
		}
	}
	item.EducationPublicDocumentType.PublicDocumentTypeID = item.ID
	return nil
}

func (item *PublicDocumentType) SetIdForChildren() {
	for i := range item.DocumentTypes {
		item.DocumentTypes[i].PublicDocumentTypeID = item.ID
	}
}
