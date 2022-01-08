package meta

import (
	"fmt"
)

func (r *Repository) getCount(table *string) (res *int, err error) {
	num := 0
	query := fmt.Sprintf("SELECT COUNT (id) FROM %s", *table)
	err = r.db.QueryRow(query).Scan(&num)
	return &num, err
}
