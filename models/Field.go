package models

import (
	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helpers/uploader"
	"github.com/uptrace/bun"
)

type Field struct {
	bun.BaseModel     `bun:"fields,alias:fields"`
	ID                uuid.UUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	Name              string    `json:"name"`
	Code              string    `json:"code"`
	Order             uint      `bun:"field_order" json:"order"`
	Comment           string    `json:"comment"`
	Required          bool      `json:"required"`
	RequiredForCancel bool      `json:"requiredForCancel"`
	Mask              string    `json:"mask"`

	Form   *Form         `bun:"rel:belongs-to" json:"form"`
	FormID uuid.NullUUID `bun:"type:uuid" json:"formId"`

	FormPattern   *FormPattern  `bun:"rel:belongs-to" json:"formPattern"`
	FormPatternID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"formPatternId"`

	ValueType   *ValueType `bun:"rel:belongs-to" json:"valueType"`
	ValueTypeID uuid.UUID  `bun:"type:uuid" json:"valueTypeId"`

	File   *FileInfo     `bun:"rel:belongs-to" json:"file"`
	FileID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"fileId"`

	FormValue   *FormValue    `bun:"rel:belongs-to" json:"formValue"`
	FormValueID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"formValueId"`

	MaskTokens          MaskTokens  `bun:"rel:has-many" json:"maskTokens"`
	MaskTokensForDelete []uuid.UUID `bun:"-" json:"maskTokensForDelete"`
}

type Fields []*Field

func (item *Field) SetIDForChildren() {
	for i := range item.MaskTokens {
		item.MaskTokens[i].FieldID = item.ID
	}
}

func (items Fields) SetIDForChildren() {
	for i := range items {
		items[i].SetIDForChildren()
	}
}

func (item *Field) GetMaskTokens() MaskTokens {
	items := make(MaskTokens, 0)
	items = append(items, item.MaskTokens...)
	return items
}

func (items Fields) GetMaskTokens() MaskTokens {
	itemsForGet := make(MaskTokens, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].GetMaskTokens()...)
	}
	return itemsForGet
}

func (items Fields) GetMaskTokensForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].MaskTokensForDelete...)
	}
	return itemsForGet
}

func (item *Field) SetForeignKeys() {
	item.ValueTypeID = item.ValueType.ID
	item.FileID = item.File.ID
	item.FormID = item.Form.ID
	item.FormPatternID = item.FormPattern.ID
	item.FormValueID = item.FormValue.ID
}

func (item *Field) SetFilePath(fileID *string) *string {
	if item.File.ID.UUID.String() == *fileID {
		item.File.FileSystemPath = uploader.BuildPath(fileID)
		return &item.File.FileSystemPath
	}
	return nil
}

func (item *Field) GetFileInfos() FileInfos {
	items := make(FileInfos, 0)
	if item.File.FileSystemPath != "" && item.File.OriginalName != "" {
		items = append(items, item.File)
	}
	return items
}

func (items Fields) SetFilePath(fileID string) *string {
	for i := range items {
		if items[i].File.ID.UUID.String() == fileID {
			items[i].File.FileSystemPath = uploader.BuildPath(&fileID)
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
			if items[i].File.FileSystemPath != "" && items[i].File.OriginalName != "" {
				items[i].FileID = items[i].File.ID
			}
		}
		if items[i].Form != nil {
			items[i].FormID = items[i].Form.ID
		}
		if items[i].FormPattern != nil {
			items[i].FormPatternID = items[i].FormPattern.ID
		}
	}
}
