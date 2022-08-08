package formstatusgroups

import (
	"fmt"
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) getAll() (item models.FormStatusGroupsWithCount, err error) {
	query := r.db().NewSelect().
		Model(&item.FormStatusGroups).
		Relation("FormStatuses")
	r.queryFilter.HandleQuery(query)
	err = query.Scan(r.ctx)
	//err = query.Scan(r.ctx)
	fmt.Println(item.FormStatusGroups)
	fmt.Println(item.FormStatusGroups)
	fmt.Println(item.FormStatusGroups)
	fmt.Println(item.FormStatusGroups)
	fmt.Println(item.FormStatusGroups)
	return item, err
}

func (r *Repository) get(id *string) (*models.FormStatusGroup, error) {
	item := models.FormStatusGroup{}
	err := r.db().NewSelect().Model(&item).
		Relation("FormStatuses").
		Where("form_status_groups.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) upsert(item *models.FormStatusGroup) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(item).
		Set("id = EXCLUDED.id").
		Set("name = EXCLUDED.name").
		Set("code = EXCLUDED.code").
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.FormStatusGroups) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("name = EXCLUDED.name").
		Set("code = EXCLUDED.code").
		Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.FormStatusGroup{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}
