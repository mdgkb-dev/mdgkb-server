package roles

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	r.queryFilter, err = r.helper.SQL.CreateQueryFilter(c)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) getAll() (models.Roles, error) {
	items := make(models.Roles, 0)
	query := r.db.NewSelect().
		Model(&items)
	r.queryFilter.Paginator.CreatePagination(query)
	r.queryFilter.Filter.CreateFilter(query)
	r.queryFilter.Sorter.CreateOrder(query)
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Role, error) {
	item := models.Role{}
	err := r.db.NewSelect().Model(&item).
		Where("roles.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) getDefaultRole() (*models.Role, error) {
	item := models.Role{}
	err := r.db.NewSelect().Model(&item).
		Where("roles.name = ?", models.RoleNameUser).Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.Role) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Role{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Role) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
