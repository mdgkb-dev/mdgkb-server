package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type AdmissionCommitteeDocumentType struct {
	bun.BaseModel  `bun:"admission_committee_document_types,alias:admission_committee_document_types"`
	ID             uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" json:"id,omitempty"`
	Order          int           `bun:"admission_committee_document_type_order" json:"order"`
	DocumentTypeID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"documentTypeId"`
	DocumentType   *DocumentType `bun:"rel:belongs-to" json:"documentType"`
}

type AdmissionCommitteeDocumentTypes []*AdmissionCommitteeDocumentType

func (item *AdmissionCommitteeDocumentType) SetForeignKeys() {
	item.DocumentTypeID = item.DocumentType.ID
}

func (items AdmissionCommitteeDocumentTypes) SetForeignKeys() {
	for i := range items {
		items[i].SetForeignKeys()
	}
}

func (item *AdmissionCommitteeDocumentType) SetFilePath(fileID *string) *string {
	path := item.DocumentType.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}

func (items AdmissionCommitteeDocumentTypes) SetFilePath(fileID *string) *string {
	for i := range items {
		path := items[i].SetFilePath(fileID)
		if path != nil {
			return path
		}
	}
	return nil
}

func (items AdmissionCommitteeDocumentTypes) GetDocumentTypes() DocumentTypes {
	itemsForGet := make(DocumentTypes, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.DocumentType)
	}
	return itemsForGet
}
