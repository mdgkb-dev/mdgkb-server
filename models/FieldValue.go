package models

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type FieldValue struct {
	bun.BaseModel `bun:"field_values,alias:field_values"`
	ID            uuid.UUID  `bun:"type:uuid,default:uuid_generate_v4()" json:"id" `
	ValueString   string     `json:"valueString"`
	ValueNumber   int        `json:"valueNumber"`
	ValueDate     *time.Time `json:"valueDate"`

	Field   *Field    `bun:"rel:belongs-to" json:"field"`
	FieldID uuid.UUID `bun:"type:uuid" json:"fieldId"`

	EventApplication   *EventApplication `bun:"rel:belongs-to" json:"eventApplication"`
	EventApplicationID uuid.UUID         `bun:"type:uuid" json:"eventApplicationId"`

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
