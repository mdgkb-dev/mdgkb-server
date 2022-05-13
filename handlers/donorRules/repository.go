package donorRules

import (
	"github.com/google/uuid"
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) getAll(userID *uuid.UUID) (models.DonorRules, error) {
	items := make(models.DonorRules, 0)
	q := r.db.NewSelect().
		Model(&items).
		Relation("Image")
	if userID != nil {
		q = q.Relation("DonorRulesUsers", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.Where("user_id = ?", userID)
		})
	}
	err := q.Order("donor_rules.donor_rule_order").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.DonorRule)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.DonorRules) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("name = EXCLUDED.name").
		Set("image_id = EXCLUDED.image_id").
		Set("donor_rule_order = EXCLUDED.donor_rule_order").
		Exec(r.ctx)
	return err
}

func (r *Repository) addToUser(item *models.DonorRuleUser) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) deleteFromUser(item *models.DonorRuleUser) (err error) {
	_, err = r.db.NewDelete().Model(item).
		Where("user_id = ? and donor_rule_id = ?", item.UserID, item.DonorRuleID).Exec(r.ctx)
	return err
}
