package educationalOrganization

import (
	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}
