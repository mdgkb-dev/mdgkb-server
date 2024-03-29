package models

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/helpers/uploader"
	"github.com/uptrace/bun"
)

type FieldValue struct {
	bun.BaseModel `bun:"field_values,alias:field_values"`
	ID            uuid.UUID  `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id" `
	ValueString   string     `json:"valueString"`
	ValueNumber   int        `json:"valueNumber"`
	ValueDate     *time.Time `bun:",nullzero" json:"valueDate"`
	ModChecked    bool       `json:"modChecked"`
	ModComment    string     `json:"modComment"`

	Field   *Field    `bun:"rel:belongs-to" json:"field"`
	FieldID uuid.UUID `bun:"type:uuid" json:"fieldId"`

	FormValue   *FormValue    `bun:"rel:belongs-to" json:"formValue"`
	FormValueID uuid.NullUUID `bun:"type:uuid,nullzero,default:NULL" json:"formValueId"`

	EventApplication   *EventApplication `bun:"rel:belongs-to" json:"eventApplication"`
	EventApplicationID uuid.NullUUID     `bun:"type:uuid,nullzero,default:NULL" json:"eventApplicationId"`

	File   *FileInfo     `bun:"rel:belongs-to" json:"file"`
	FileID uuid.NullUUID `json:"fileId"`

	FieldValuesFiles          FieldValuesFiles `bun:"rel:has-many" json:"fieldValuesFiles"`
	FieldValuesFilesForDelete []uuid.UUID      `bun:"-" json:"fieldValuesFilesForDelete"`

	Value string `bun:"-" json:"-"`
}

type FieldValues []*FieldValue

func (items FieldValues) sortByFieldName() {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Field.Order < items[j].Field.Order
	})
}

func (items FieldValues) PrepareValuesForPrint() {
	items.sortByFieldName()
	items.writeValueToPrint()
}

func (items FieldValues) writeValueToPrint() {
	for i := range items {
		items[i].writeValueToPrint()
	}
}

func (item *FieldValue) writeValueToPrint() {
	if item.ValueDate != nil {
		item.Value = fmt.Sprintf("%d-%d-%d", item.ValueDate.Year(), item.ValueDate.Month(), item.ValueDate.Day())
	}
	if item.ValueString != "" {
		item.Value = item.ValueString
	}
	if item.ValueNumber != 0 {
		item.Value = strconv.Itoa(item.ValueNumber)
	}
}

func (item *FieldValue) GetFileInfos() FileInfos {
	items := make(FileInfos, 0)
	items = append(items, item.File)
	return items
}

func (items FieldValues) GetFieldValuesFiles() FieldValuesFiles {
	itemsForGet := make(FieldValuesFiles, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].GetFieldValuesFiles()...)
	}
	return itemsForGet
}

func (items FieldValues) GetFieldValuesFilesForDelete() []uuid.UUID {
	itemsForGet := make([]uuid.UUID, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].FieldValuesFilesForDelete...)
	}
	return itemsForGet
}

func (item *FieldValue) GetFieldValuesFiles() FieldValuesFiles {
	items := make(FieldValuesFiles, 0)
	items = append(items, item.FieldValuesFiles...)
	return items
}

func (item *FieldValue) SetFilePath(fileID *string) *string {
	if item.File.ID.UUID.String() == *fileID {
		item.File.FileSystemPath = uploader.BuildPath(fileID)
		return &item.File.FileSystemPath
	}
	return nil
}

func (items FieldValues) SetFilePath(fileID string) *string {
	for i := range items {
		if items[i].File.ID.UUID.String() == fileID {
			items[i].File.FileSystemPath = uploader.BuildPath(&fileID)
			return &items[i].File.FileSystemPath
		}
	}
	return nil
}

func (items FieldValues) GetFileInfos() FileInfos {
	itemsForGet := make(FileInfos, 0)
	for i := range items {
		itemsForGet = append(itemsForGet, items[i].File)
	}
	return itemsForGet
}

func (items FieldValues) SetForeignKeys() {
	for i := range items {
		if items[i].File != nil {
			items[i].FileID = items[i].File.ID
		}
		if items[i].EventApplication != nil {
			items[i].EventApplicationID = items[i].EventApplication.ID
		}
		if items[i].Field != nil {
			items[i].FieldID = items[i].Field.ID
		}
		fmt.Println(items[i].FileID)
	}
}

func (items FieldValues) GetFields() Fields {
	itemsForGet := make(Fields, 0)
	for _, item := range items {
		if item.Field != nil {
			itemsForGet = append(itemsForGet, item.Field)
		}
	}
	return itemsForGet
}

func (item *FieldValue) SetIDForChildren() {
	for i := range item.FieldValuesFiles {
		item.FieldValuesFiles[i].FieldValueID = item.ID
	}
}

func (items FieldValues) SetIDForChildren() {
	for i := range items {
		items[i].SetIDForChildren()
	}
}
