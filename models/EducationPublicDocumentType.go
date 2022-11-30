package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type EducationPublicDocumentType struct {
	bun.BaseModel        `bun:"education_public_document_types,alias:education_public_document_types"`
	ID                   uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id,omitempty"`
	PublicDocumentTypeID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"publicDocumentTypeId"`
	PublicDocumentType   *PageSideMenu `bun:"rel:belongs-to" json:"publicDocumentType"`
}

type EducationPublicDocumentTypes []*EducationPublicDocumentType

func (item *EducationPublicDocumentType) SetForeignKeys() {
	item.PublicDocumentTypeID = item.PublicDocumentType.ID
}

func (items EducationPublicDocumentTypes) SetForeignKeys() {
	for i := range items {
		items[i].SetForeignKeys()
	}
}

func (item *EducationPublicDocumentType) SetFilePath(fileID *string) *string {
	path := item.PublicDocumentType.SetFilePath(fileID)
	if path != nil {
		return path
	}
	return nil
}

func (items EducationPublicDocumentTypes) SetFilePath(fileID *string) *string {
	for i := range items {
		path := items[i].SetFilePath(fileID)
		if path != nil {
			return path
		}
	}
	return nil
}
