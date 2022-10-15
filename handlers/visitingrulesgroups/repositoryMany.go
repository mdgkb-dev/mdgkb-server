package visitingrulesgroups

import (
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items models.VisitingRulesGroups) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.VisitingRuleGroup)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.VisitingRulesGroups) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Set("name = EXCLUDED.name").
		Set("color = EXCLUDED.color").
		Set("visiting_rule_group_order = EXCLUDED.visiting_rule_group_order").
		Set("division_id = EXCLUDED.division_id").
		Model(&items).
		Exec(r.ctx)
	return err
}
