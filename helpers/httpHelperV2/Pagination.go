package httpHelper

import (
	"encoding/json"
	"github.com/uptrace/bun"
)

type Pagination struct {
	Offset     *int    `json:"offset"`
	Limit      *int    `json:"limit"`
	CursorMode bool    `json:"cursorMode"`
	Cursor     *Cursor `json:"cursor"`
}

// ParseJSONToPagination constructor
func parseJSONToPagination(args string) (pagination *Pagination, err error) {
	err = json.Unmarshal([]byte(args), &pagination)
	if err != nil {
		return pagination, err
	}
	return pagination, err
}

func (p *Pagination) CreatePagination(query *bun.SelectQuery) {
	if p.CursorMode {
		p.Cursor.createPagination(query)
	} else {
		query = query.Offset(*p.Offset)
	}
	query = query.Limit(*p.Limit)
}
