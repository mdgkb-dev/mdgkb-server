package formPatterns

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

func (r *Repository) getAll() (models.FormPatterns, error) {
	items := make(models.FormPatterns, 0)
	query := r.db.NewSelect().
		Model(&items).
		Relation("Fields.File").
		Relation("Fields.ValueType")

	r.queryFilter.HandleQuery(query)
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.FormPattern, error) {
	item := models.FormPattern{}
	err := r.db.NewSelect().Model(&item).
		Relation("Fields", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("fields.field_order")
		}).
		Relation("Fields.File").
		Relation("Fields.ValueType").
		Relation("FormStatusGroup.FormStatuses").
		Relation("DefaultFormStatus").
		Relation("PersonalDataAgreement").
		Relation("Fields.MaskTokens").
		Where("form_patterns.id = ?", id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.FormPattern) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.FormPattern{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.FormPattern) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
