package visitingRules

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.VisitingRule) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.VisitingRules, error) {
	items := make(models.VisitingRules, 0)
	err := r.db.NewSelect().Model(&items).Order("rule_order").Where("division_id is null").Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.VisitingRule, error) {
	item := models.VisitingRule{}
	err := r.db.NewSelect().Model(&item).Where("id = ?", id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db.NewDelete().Model(&models.VisitingRule{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.VisitingRules) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Set("text = EXCLUDED.text").
		Set("rule_order = EXCLUDED.rule_order").
		Set("is_list_item = EXCLUDED.is_list_item").
		Model(&items).
		Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.VisitingRule)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.VisitingRule) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
