package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/helpers/uploadHelper"
)

type EducationalOrganizationDocumentTypeDocument struct {
	bun.BaseModel `bun:"educational_organization_document_types_documents,alias:educational_organization_document_types_documents"`
	ID   uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id"`

	EducationalOrganizationDocumentType *EducationalOrganizationDocumentType `bun:"rel:belongs-to" json:"educationalOrganizationDocumentType"`
	EducationalOrganizationDocumentTypeID uuid.UUID `bun:"type:uuid" json:"educationalOrganizationDocumentTypeId"`

	Document *Document `bun:"rel:belongs-to" json:"document"`
	DocumentID uuid.UUID `bun:"type:uuid" json:"documentId"`
}

type EducationalOrganizationDocumentTypeDocuments []*EducationalOrganizationDocumentTypeDocument

func (i EducationalOrganizationDocumentTypeDocuments) GetDocuments() Documents {
	items := make(Documents, 0)
	for _, item := range i {
		items = append(items, item.Document)
	}
	return items
}

func (items EducationalOrganizationDocumentTypeDocuments) SetForeignKeys()  {
	for i := range items {
		items[i].DocumentID = items[i].Document.ID
	}
}

func (i EducationalOrganizationDocumentTypeDocuments) SetFilePath(fileId *string) *string {
	for _, item := range i {
		if item.Document.FileInfo.ID.UUID.String() == *fileId {
			item.Document.FileInfo.FileSystemPath = uploadHelper.BuildPath(fileId)
			return &item.Document.FileInfo.FileSystemPath
		}
	}
	return nil
}
