package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type HospitalizationToDocumentType struct {
	bun.BaseModel `bun:"hospitalizations_to_document_types,alias:vacancies_responses_to_document_types"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id,omitempty"`

	DocumentType   *DocumentType `bun:"rel:belongs-to" json:"documentType"`
	DocumentTypeID uuid.UUID     `bun:"type:uuid"  json:"documentTypeId"`

	Hospitalization   *Hospitalization `bun:"rel:belongs-to" json:"hospitalization"`
	HospitalizationID uuid.UUID        `bun:"type:uuid"  json:"hospitalizationId"`
}

type HospitalizationsToDocumentTypes []*HospitalizationToDocumentType

// func (item *VacancyResponseToDocument) SetForeignKeys() {
// 	item.DocumentID = item.Document.ID
// }

// func (items VacancyResponsesToDocuments) SetForeignKeys() {
// 	for i := range items {
// 		items[i].SetForeignKeys()
// 	}
// }

// func (items VacancyResponsesToDocuments) GetDocuments() Documents {
// 	itemsForGet := make(Documents, len(items))
// 	for i := range items {
// 		itemsForGet[i] = items[i].Document
// 	}
// 	return itemsForGet
// }

// func (item *VacancyResponseToDocument) SetFilePath(fileID *string) *string {
// 	return item.Document.SetFilePath(fileID)
// }

// func (items VacancyResponsesToDocuments) SetFilePath(fileID *string) *string {
// 	for i := range items {
// 		return items[i].SetFilePath(fileID)
// 	}
// 	return nil
// }
