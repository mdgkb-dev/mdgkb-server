package dailymenuorders

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

func (r *Repository) create(item *models.DailyMenuOrder) (err error) {
	_, err = r.db().NewInsert().Model(item).ExcludeColumn("user_id", "number").Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (item models.DailyMenuOrdersWithCount, err error) {
	item.DailyMenuOrders = make(models.DailyMenuOrders, 0)
	query := r.db().NewSelect().Model(&item.DailyMenuOrders).
		Relation("DailyMenuOrderItems.DailyMenuItem.DishSample").
		Relation("FormValue.Child.Human").
		Relation("FormValue.User.Human").
		Relation("FormValue.Fields.ValueType").
		Relation("FormValue.FieldValues.File").
		Relation("FormValue.FieldValues.Field.ValueType").
		Relation("FormValue.FormStatus.FormStatusToFormStatuses")
	item.Count, err = query.ScanAndCount(r.ctx)
	return item, err
}

func (r *Repository) get(id string) (*models.DailyMenuOrder, error) {
	item := models.DailyMenuOrder{}
	err := r.db().NewSelect().Model(&item).
		Relation("DailyMenuOrderItems.DailyMenuItem.DishSample").
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
		Where("?TableAlias.id = ?", id).
		Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.DailyMenuOrder{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.DailyMenuOrder) (err error) {
	_, err = r.db().NewUpdate().Model(item).ExcludeColumn("user_id", "number").Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.DailyMenuOrders) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		ExcludeColumn("user_id", "number").
		Exec(r.ctx)
	return err
}
