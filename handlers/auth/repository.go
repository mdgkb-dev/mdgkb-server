package auth

import (
	_ "github.com/go-pg/pg/v10/orm"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/middleware"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) saveClientPermissions(paths []string) (err error) {
	cas := make([]*middleware.CasbinRule, 0)
	for _, path := range paths {
		rule := middleware.CasbinRule{Ptype: "p", V1: path}
		cas = append(cas, &rule)
	}
	_, err = r.db.NewInsert().
		Model(&cas).
		Exec(r.ctx)
	return err
}
