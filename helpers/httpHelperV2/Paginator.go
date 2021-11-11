package httpHelper

import "github.com/uptrace/bun"

func CreatePaginationQuery(query *bun.SelectQuery, pagination *Pagination) {
	query = query.Offset(*pagination.Offset)
	query = query.Limit(*pagination.Limit)
}
