package httpHelper

import (
	"fmt"
	"github.com/uptrace/bun"
)

type Cursor struct {
	Operator Operator `json:"operation"`
	Column   string   `json:"column"`
	Value    string   `json:"value"`
	Initial  bool     `json:"initial"`
}

func (c *Cursor) createPagination(query *bun.SelectQuery) {
	if c.Initial {
		return
	}
	q := fmt.Sprintf("%s %s '%s'", c.Column, c.Operator, c.Value)
	query = query.Where(q)
}
