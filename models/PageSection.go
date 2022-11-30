package models

import (
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"github.com/uptrace/bun"
)

type PageSection struct {
	bun.BaseModel `bun:"page_sections,alias:page_sections"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id,omitempty"`
	Name          string        `json:"name,omitempty"`
	Order         uint          `bun:"document_type_order" json:"order"`
	Description   string        `json:"description,omitempty"`

	PageSideMenuID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"pageSideMenuId"`
	PageSideMenu   *PageSideMenu `bun:"rel:belongs-to" json:"pageSideMenu"`

	PageSectionDocuments          PageSectionDocuments `bun:"rel:has-many" json:"pageSectionDocuments"`
	PageSectionDocumentsForDelete []uuid.UUID          `bun:"-" json:"pageSectionDocumentsForDelete"`

	PageSectionImagesForDelete []uuid.UUID       `bun:"-" json:"pageSectionImagesForDelete"`
	PageSectionImages          PageSectionImages `bun:"rel:has-many" json:"pageSectionImages"`
}

type PageSections []*PageSection

func (item *PageSection) SetIDForChildren() {
	for i := range item.PageSectionDocuments {
		item.PageSectionDocuments[i].PageSectionID = item.ID
	}
	for i := range item.PageSectionImages {
		item.PageSectionImages[i].DocumentTypeID = item.ID
	}
}

func (items PageSections) SetIDForChildren() {
	for i := range items {
		items[i].SetIDForChildren()
	}
}

func (item PageSection) SetFilePath(fileID *string) *string {
	for i := range item.PageSectionDocuments {
		filePath := item.PageSectionDocuments[i].SetFilePath(fileID)
		if filePath != nil {
			return filePath
		}
	}
	for i := range item.PageSectionImages {
		if item.PageSectionImages[i].FileInfo.ID.UUID.String() == *fileID {
			item.PageSectionImages[i].FileInfo.FileSystemPath = uploadHelper.BuildPath(fileID)
			return &item.PageSectionImages[i].FileInfo.FileSystemPath
		}
	}
	return nil
}

func (items PageSections) GetDocuments() PageSectionDocuments {
	itemsForGet := make(PageSectionDocuments, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PageSectionDocuments...)
	}
	return itemsForGet
}

func (items PageSections) GetDocumentsIDForDelete() []uuid.UUID {
	idPool := make([]uuid.UUID, 0)
	for _, item := range items {
		idPool = append(idPool, item.PageSectionDocumentsForDelete...)
	}
	return idPool
}

func (items PageSections) GetDocumentTypeImages() PageSectionImages {
	itemsForGet := make(PageSectionImages, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PageSectionImages...)
	}
	return itemsForGet
}

func (items PageSections) GetDocumentTypeImagesForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.PageSectionImagesForDelete...)
	}
	return itemsForGet
}

// func (items PageSections) GetFileInfos() FileInfos {
// 	itemsForGet := make(FileInfos, 0)
// 	for _, item := range items {
// 		itemsForGet = append(itemsForGet, item.Scans...)
// 		itemsForGet = append(itemsForGet, item.Scan)
// 	}
// 	return itemsForGet
// }

// func (items PageSections) SetFileInfoID() {
// 	for _, item := range items {
// 		item.ScanID = item.Scan.ID
// 	}
// }
