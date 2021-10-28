package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type EducationalOrganizationDocumentType struct {
	bun.BaseModel `bun:"educational_organization_document_types,alias:educational_organization_document_types"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`
	Name          string    `json:"name"`

	EducationalOrganizationDocumentTypeDocuments          EducationalOrganizationDocumentTypeDocuments `bun:"rel:has-many" json:"educationalOrganizationDocumentTypeDocuments"`
	EducationalOrganizationDocumentTypeDocumentsForDelete []string                                     `bun:"-" json:"educationalOrganizationDocumentTypeDocumentsForDelete"`
}

type EducationalOrganizationDocumentTypes []*EducationalOrganizationDocumentType

func (i EducationalOrganizationDocumentTypes) GetIDForDelete() []string {
	idPool := make([]string, 0)
	for _, item := range i {
		idPool = append(idPool, item.EducationalOrganizationDocumentTypeDocumentsForDelete...)
	}
	return idPool
}

func (i EducationalOrganizationDocumentTypes) GetEducationalOrganizationDocumentTypeDocuments() EducationalOrganizationDocumentTypeDocuments {
	items := make(EducationalOrganizationDocumentTypeDocuments, 0)
	for _, item := range i {
		items = append(items, item.EducationalOrganizationDocumentTypeDocuments...)
	}
	return items
}

func (i EducationalOrganizationDocumentTypes) SetChildrenForeignKeys() {
	for id := range i {
		i[id].SetChildrenForeignKeys()
	}
}

func (i *EducationalOrganizationDocumentType) SetChildrenForeignKeys() {
	for index := range i.EducationalOrganizationDocumentTypeDocuments {
		i.EducationalOrganizationDocumentTypeDocuments[index].EducationalOrganizationDocumentTypeID = i.ID
	}
}

func (i EducationalOrganizationDocumentTypes) SetFilePath(fileID *string) *string {
	for _, item := range i {
		path := item.EducationalOrganizationDocumentTypeDocuments.SetFilePath(fileID)
		if path != nil {
			return path
		}
	}
	return nil
}
