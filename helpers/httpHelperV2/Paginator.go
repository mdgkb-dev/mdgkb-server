package httpHelper

import "github.com/uptrace/bun"

func (i *HTTPHelper) CreatePaginationQuery(query *bun.SelectQuery, pagination *Pagination) {
	if pagination != nil {
		query = query.Limit(*pagination.Limit)
		query = query.Offset(*pagination.Offset)
	}
}
