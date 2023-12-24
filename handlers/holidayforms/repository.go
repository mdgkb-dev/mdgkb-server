package holidayforms

import (
	"fmt"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.HolidayForm) error {
	fmt.Println(item)
	_, err := r.db().NewInsert().
		Model(item).
		// Set("needing = ?", pgdialect.Array(item.Needing)).
		// Set("place = ?", pgdialect.Array(item.Place)).
		Exec(r.ctx)

	return err
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}
