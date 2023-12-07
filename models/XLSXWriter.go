package models

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type XLSXWriter struct {
	bun.BaseModel `bun:"data_queries,alias:data_queries"`
	ID            uuid.NullUUID `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name          string        `json:"name"`
	Type          string        `bun:"type:register_query_type_enum" json:"type"`

	WithAge         bool `json:"withAge"`
	CountAverageAge bool `json:"countAverageAge"`
}

func (item *XLSXWriter) setStyle() {
	// height := 6 + len(item.ResearchesPool.RegisterToPatient)
	// xl.SetBorder(height)
	// item.ResearchQueryGroups.writeAggregates(xl)
	// xl.AutofitAllColumns()
}

func (item *XLSXWriter) WriteFile(headers [][]interface{}, agregator Agregator, data [][]interface{}) ([]byte, error) {
	return []byte{}, nil
}
