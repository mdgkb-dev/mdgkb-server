package gates

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) setQueryFilter(c *gin.Context) (err error) {
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) getAll() (models.Gates, error) {
	items := make(models.Gates, 0)
	query := r.db().NewSelect().
		Model(&items).
		Relation("FormPattern")
	//Relation("VisitsApplication.Division")
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.Gate, error) {
	item := models.Gate{}
	err := r.db().NewSelect().Model(&item).
		Relation("FormPattern.Fields", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("fields.field_order")
		}).
		Relation("FormPattern.Fields.File").
		Relation("FormPattern.DefaultFormStatus").
		Relation("FormPattern.FormStatusGroup").
		Relation("FormPattern.Fields.ValueType").
		Relation("FormPattern.PersonalDataAgreement").
		Where("gates.id = ?", id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.Gate) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Gate{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Gate) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Gates) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("id = EXCLUDED.id").
		Set("name = EXCLUDED.name").
		Set("form_pattern_id = EXCLUDED.form_pattern_id").
		Exec(r.ctx)
	return err
}
