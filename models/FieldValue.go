package models

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/pro-assistance/pro-assister/uploadHelper"
	"github.com/uptrace/bun"
)

type FieldValue struct {
	bun.BaseModel `bun:"field_values,alias:field_values"`
	ID            uuid.UUID  `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	ValueString   string     `json:"valueString"`
	ValueNumber   int        `json:"valueNumber"`
	ValueDate     *time.Time `bun:",nullzero" json:"valueDate"`
	ModChecked    bool       `json:"modChecked"`

	Field   *Field    `bun:"rel:belongs-to" json:"field"`
	FieldID uuid.UUID `bun:"type:uuid" json:"fieldId"`

	EventApplication   *EventApplication `bun:"rel:belongs-to" json:"eventApplication"`
	EventApplicationID uuid.NullUUID     `bun:"type:uuid,nullzero,default:NULL" json:"eventApplicationId"`

	DpoApplication   *DpoApplication `bun:"rel:belongs-to" json:"dpoApplication"`
	DpoApplicationID uuid.NullUUID   `bun:"type:uuid,nullzero,default:NULL" json:"dpoApplicationId"`

	// PostgraduateApplication   *PostgraduateApplication `bun:"rel:belongs-to" json:"postgraduateApplication"`
	// PostgraduateApplicationID uuid.NullUUID            `bun:"type:uuid,nullzero,default:NULL" json:"postgraduateApplicationId"`

	File   *FileInfo     `bun:"rel:belongs-to" json:"file"`
	FileID uuid.NullUUID `json:"fileId"`

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

func (i *FieldValue) writeValueToPrint() {
	if i.ValueDate != nil {
		i.Value = fmt.Sprintf("%d-%d-%d", i.ValueDate.Year(), i.ValueDate.Month(), i.ValueDate.Day())
	}
	if i.ValueString != "" {
		i.Value = i.ValueString
	}
	if i.ValueNumber != 0 {
		i.Value = strconv.Itoa(i.ValueNumber)
	}
}

func (item *FieldValue) GetFileInfos() FileInfos {
	items := make(FileInfos, 0)
	items = append(items, item.File)
	return items
}

func (item *FieldValue) SetFilePath(fileID *string) *string {
	if item.File.ID.UUID.String() == *fileID {
		item.File.FileSystemPath = uploadHelper.BuildPath(fileID)
		return &item.File.FileSystemPath
	}
	return nil
}

func (items FieldValues) SetFilePath(fileID string) *string {
	for i := range items {
		if items[i].File.ID.UUID.String() == fileID {
			items[i].File.FileSystemPath = uploadHelper.BuildPath(&fileID)
			return &items[i].File.FileSystemPath
		}
	}
	return nil
}

func (items FieldValues) GetFileInfos() FileInfos {
	itemsForGet := make(FileInfos, 0)
	for _, item := range items {
		itemsForGet = append(itemsForGet, item.File)
	}
	return itemsForGet
}

func (items FieldValues) SetForeignKeys() {
	for i := range items {
		if items[i].File != nil {
			items[i].FileID = items[i].File.ID
		}
		if items[i].DpoApplication != nil {
			items[i].DpoApplicationID = items[i].DpoApplication.ID
		}
		if items[i].EventApplication != nil {
			items[i].EventApplicationID = items[i].EventApplication.ID
		}
		if items[i].Field != nil {
			items[i].FieldID = items[i].Field.ID
		}
	}
}
