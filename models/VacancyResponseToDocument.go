package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type VacancyResponseToDocument struct {
	bun.BaseModel `bun:"vacancies_responses_to_documents,alias:vacancies_responses_to_documents"`
	ID            uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()" json:"id" json:"id,omitempty"`

	Document   *Document `bun:"rel:belongs-to" json:"document"`
	DocumentID uuid.UUID `bun:"type:uuid"  json:"documentId"`

	VacancyResponse   *VacancyResponse `bun:"rel:belongs-to" json:"vacancyResponse"`
	VacancyResponseID uuid.UUID        `bun:"type:uuid"  json:"vacancyResponseId"`
}

type VacancyResponsesToDocuments []*VacancyResponseToDocument

//
func (item *VacancyResponseToDocument) SetForeignKeys() {
	item.DocumentID = item.Document.ID
}

func (items VacancyResponsesToDocuments) SetForeignKeys() {
	for i := range items {
		items[i].SetForeignKeys()
	}
}

//
//func (items Documents) SetIdForChildren() {
//	for i := range items {
//		items[i].SetIdForChildren()
//	}
//}
//
func (items VacancyResponsesToDocuments) GetDocuments() Documents {
	itemsForGet := make(Documents, len(items))
	for i := range items {
		itemsForGet[i] = items[i].Document
	}
	return itemsForGet
}

//
//func (items Documents) SetFileInfoID() {
//	for _, item := range items {
//		item.ScanID = item.Scan.ID
//	}
//}
//

func (item *VacancyResponseToDocument) SetFilePath(fileID *string) *string {
	return item.Document.SetFilePath(fileID)
}

func (items VacancyResponsesToDocuments) SetFilePath(fileID *string) *string {
	for i := range items {
		return items[i].SetFilePath(fileID)
	}
	return nil
}

//
//func (items Documents) SetFilePath(fileID *string) *string {
//	for i := range items {
//		items[i].SetFilePath(fileID)
//	}
//	return nil
//}
