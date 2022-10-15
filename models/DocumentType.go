package models

import (
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"github.com/uptrace/bun"
)

type DocumentType struct {
	bun.BaseModel        `bun:"document_types,alias:document_types"`
	ID                   uuid.NullUUID       `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id,omitempty"`
	Name                 string              `json:"name,omitempty"`
	Order                uint                `bun:"document_type_order" json:"order"`
	Description          string              `json:"description,omitempty"`
	PublicDocumentTypeID uuid.NullUUID       `bun:"type:uuid,nullzero,default:NULL" json:"publicDocumentTypeId"`
	PublicDocumentType   *PublicDocumentType `bun:"rel:belongs-to" json:"publicDocumentType"`

	Documents          Documents   `bun:"rel:has-many" json:"documents"`
	DocumentsForDelete []uuid.UUID `bun:"-" json:"documentsForDelete"`

	DocumentTypeFields      DocumentTypeFields `bun:"rel:has-many" json:"documentFields"`
	DocumentFieldsForDelete []uuid.UUID        `bun:"-" json:"documentFieldsForDelete"`

	DocumentTypeImagesForDelete []uuid.UUID        `bun:"-" json:"documentTypeImagesForDelete"`
	DocumentTypeImages          DocumentTypeImages `bun:"rel:has-many" json:"documentTypeImages"`
}

type DocumentTypes []*DocumentType

func (item *DocumentType) SetIDForChildren() {
	for i := range item.DocumentTypeFields {
		item.DocumentTypeFields[i].DocumentTypeID = item.ID
	}
	for i := range item.Documents {
		item.Documents[i].DocumentTypeID = item.ID
	}
	for i := range item.DocumentTypeImages {
		item.DocumentTypeImages[i].DocumentTypeID = item.ID
	}
}

func (items DocumentTypes) SetIDForChildren() {
	for i := range items {
		items[i].SetIDForChildren()
	}
}

func (item DocumentType) SetFilePath(fileID *string) *string {
	for i := range item.Documents {
		filePath := item.Documents[i].SetFilePath(fileID)
		if filePath != nil {
			return filePath
		}
	}
	for i := range item.DocumentTypeImages {
		if item.DocumentTypeImages[i].FileInfo.ID.UUID.String() == *fileID {
			item.DocumentTypeImages[i].FileInfo.FileSystemPath = uploadHelper.BuildPath(fileID)
			return &item.DocumentTypeImages[i].FileInfo.FileSystemPath
		}
	}
	return nil
}

func (items DocumentTypes) GetDocuments() Documents {
	itemsForGet := make(Documents, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.Documents...)
	}
	return itemsForGet
}

func (items DocumentTypes) GetDocumentsIDForDelete() []uuid.UUID {
	idPool := make([]uuid.UUID, 0)
	for _, item := range items {
		idPool = append(idPool, item.DocumentsForDelete...)
	}
	return idPool
}

func (items DocumentTypes) GetDocumentTypeImages() DocumentTypeImages {
	itemsForGet := make(DocumentTypeImages, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.DocumentTypeImages...)
	}
	return itemsForGet
}

func (items DocumentTypes) GetDocumentTypeImagesForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.DocumentTypeImagesForDelete...)
	}
	return itemsForGet
}

// func (items DocumentTypes) GetFileInfos() FileInfos {
// 	itemsForGet := make(FileInfos, 0)
// 	for _, item := range items {
// 		itemsForGet = append(itemsForGet, item.Scans...)
// 		itemsForGet = append(itemsForGet, item.Scan)
// 	}
// 	return itemsForGet
// }

// func (items DocumentTypes) SetFileInfoID() {
// 	for _, item := range items {
// 		item.ScanID = item.Scan.ID
// 	}
// }
