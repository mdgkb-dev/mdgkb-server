package formpatterns

import (
	"context"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) GetAll(c context.Context) (models.FormPatterns, error) {
	items := make(models.FormPatterns, 0)
	query := r.helper.DB.IDB(c).NewSelect().
		Model(&items).
		Relation("Fields.File").
		Relation("Fields.ValueType")

	err := query.Scan(c)
	return items, err
}

func (r *Repository) Get(c context.Context) (*models.FormPattern, error) {
	item := models.FormPattern{}
	err := r.helper.DB.IDB(c).NewSelect().Model(&item).
		Relation("Fields", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Order("fields.field_order")
		}).
		Relation("Fields.File").
		Relation("Fields.ValueType").
		Relation("FormStatusGroup.FormStatuses").
		Relation("DefaultFormStatus").
		Relation("PersonalDataAgreement").
		Relation("Fields.MaskTokens").
		// Where("form_patterns.? = ?", bun.Safe(r..Col), r..Value).
		Scan(c)

	return &item, err
}

func (r *Repository) Create(c context.Context, item *models.FormPattern) (err error) {
	_, err = r.helper.DB.IDB(c).NewInsert().Model(item).Exec(c)
	return err
}

func (r *Repository) Delete(c context.Context, id string) (err error) {
	_, err = r.helper.DB.IDB(c).NewDelete().Model(&models.FormPattern{}).Where("id = ?", id).Exec(c)
	return err
}

func (r *Repository) Update(c context.Context, item *models.FormPattern) (err error) {
	_, err = r.helper.DB.IDB(c).NewUpdate().Model(item).Where("id = ?", item.ID).Exec(c)
	return err
}
