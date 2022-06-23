package models

import (
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"github.com/uptrace/bun"
)

type Field struct {
	bun.BaseModel `bun:"fields,alias:fields"`
	ID            uuid.UUID     `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name          string        `json:"name"`
	Code          string        `json:"code"`
	Order         uint          `bun:"field_order" json:"order"`
	Comment       string        `json:"comment"`
	Required      bool          `json:"required"`
	Form          *Form         `bun:"rel:belongs-to" json:"form"`
	FormID        uuid.NullUUID `bun:"type:uuid" json:"formId"`
	FormPattern   *FormPattern  `bun:"rel:belongs-to" json:"formPattern"`
	FormPatternID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"formPatternId"`
	ValueType     *ValueType    `bun:"rel:belongs-to" json:"valueType"`
	ValueTypeID   uuid.UUID     `bun:"type:uuid" json:"valueTypeId"`
	File          *FileInfo     `bun:"rel:belongs-to" json:"file"`
	FileID        uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"fileId"`
	FormValue     *FormValue    `bun:"rel:belongs-to" json:"formValue"`
	FormValueID   uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"formValueId"`
}

type Fields []*Field

func (item *Field) SetForeignKeys() {
	item.ValueTypeID = item.ValueType.ID
	item.FileID = item.File.ID
	item.FormID = item.Form.ID
	item.FormPatternID = item.FormPattern.ID
	item.FormValueID = item.FormValue.ID
}

func (item *Field) SetFilePath(fileID *string) *string {
	if item.File.ID.UUID.String() == *fileID {
		item.File.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.File.FileSystemPath
	}
	return nil
}

func (item *Field) GetFileInfos() FileInfos {
	items := make(FileInfos, 0)
	items = append(items, item.File)
	return items
}

func (items Fields) SetFilePath(fileID string) *string {
	for i := range items {
		if items[i].File.ID.UUID.String() == fileID {
			items[i].File.FileSystemPath = uploadHelper.BuildPath(&fileID)
			return &items[i].File.FileSystemPath
		}
	}
	return nil
}

func (items Fields) GetFileInfos() FileInfos {
	itemsForGet := make(FileInfos, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.File)
	}
	return itemsForGet
}

func (items Fields) SetForeignKeys() {
	for i := range items {
		if items[i].File != nil {
			items[i].FileID = items[i].File.ID
		}
		if items[i].Form != nil {
			items[i].FormID = items[i].Form.ID
		}
		if items[i].FormPattern != nil {
			items[i].FormPatternID = items[i].FormPattern.ID
		}
	}
}
