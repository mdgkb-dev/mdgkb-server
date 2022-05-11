package baseHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func (r *Repository) GetDB() *bun.DB {
	return r.db
}

func (r *Repository) SetQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}
