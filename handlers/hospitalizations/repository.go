package hospitalizations

import (
	"mdgkb/mdgkb-server/models"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	// _ "github.com/go-pg/pg/v10/orm"
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

func (r *Repository) create(item *models.Hospitalization) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.Hospitalizations, error) {
	items := make(models.Hospitalizations, 0)
	query := r.db().NewSelect().Model(&items).
		Relation("FormValue.Child.Human").
		Relation("FormValue.User.Human").
		Relation("FormValue.Fields.ValueType").
		Relation("FormValue.FieldValues.File").
		Relation("FormValue.FieldValues.Field.ValueType").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Relation("Division").
		Relation("HospitalizationType")
	err := query.Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.Hospitalization, error) {
	item := models.Hospitalization{}
	err := r.db().NewSelect().Model(&item).
		Relation("FormValue.User.Human").
		Relation("FormValue.Fields", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("fields.field_order")
		}).
		Relation("FormValue.Fields.File").
		Relation("FormValue.FormValueFiles.File").
		Relation("FormValue.Fields.ValueType").
		Relation("FormValue.FieldValues.File").
		Relation("FormValue.FieldValues.Field.ValueType").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses.ChildFormStatus").
		Relation("Division").
		Relation("HospitalizationType.FormPattern", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.ExcludeColumn("with_personal_data_agreement")
		}).
		Relation("HospitalizationType.FormPattern.Fields", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("fields.field_order")
		}).
		Relation("HospitalizationType.FormPattern.Fields.ValueType").
		Where("hospitalizations_view.id = ?", id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Hospitalization{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Hospitalization) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
