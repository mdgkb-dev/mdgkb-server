package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type EducationalOrganizationDocumentTypeDocument struct {
	bun.BaseModel `bun:"educational_organization_document_types_documents,alias:educational_organization_document_types_documents"`
	ID            uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`

	EducationalOrganizationDocumentType   *EducationalOrganizationDocumentType `bun:"rel:belongs-to" json:"educationalOrganizationDocumentType"`
	EducationalOrganizationDocumentTypeID uuid.UUID                            `bun:"type:uuid" json:"educationalOrganizationDocumentTypeId"`

	Document   *PageSectionDocument `bun:"rel:belongs-to" json:"document"`
	DocumentID uuid.UUID            `bun:"type:uuid" json:"documentId"`
}

type EducationalOrganizationDocumentTypeDocuments []*EducationalOrganizationDocumentTypeDocument

func (items EducationalOrganizationDocumentTypeDocuments) GetDocuments() PageSectionDocuments {
	itemsForGet := make(PageSectionDocuments, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.Document)
	}
	return itemsForGet
}

func (items EducationalOrganizationDocumentTypeDocuments) SetForeignKeys() {
	for i := range items {
		items[i].DocumentID = items[i].Document.ID
	}
}

func (items EducationalOrganizationDocumentTypeDocuments) SetFilePath(fileID *string) *string {
	for range items {
		//if item.PageSectionDocument.Scan.ID.UUID.String() == *fileID {
		//	item.PageSectionDocument.Scan.FileSystemPath = uploadHelper.BuildPath(fileID)
		//	return &item.PageSectionDocument.Scan.FileSystemPath
		//}
	}
	return nil
}
