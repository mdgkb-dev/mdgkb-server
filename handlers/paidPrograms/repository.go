package paidPrograms

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) createMany(items models.PaidPrograms) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.PaidProgram)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.PaidPrograms) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Set("name = EXCLUDED.name").
		Set("paid_programs_group_id = EXCLUDED.paid_programs_group_id").
		Model(&items).
		Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.PaidProgram) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) get(id string) (*models.PaidProgram, error) {
	item := models.PaidProgram{}
	err := r.db.NewSelect().
		Model(&item).
		Relation("PaidProgramPackages.PaidProgramServicesGroups.PaidProgramServices").
		Relation("PaidProgramPackages.PaidProgramPackagesOptions").
		Relation("PaidProgramOptionsGroups", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Order("item_order")
		}).
		Relation("PaidProgramOptionsGroups.PaidProgramOptions", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Order("item_order")
		}).
		Relation("PaidProgramsGroup").
		Where("paid_programs.id = ?", id).
		Scan(r.ctx)

	return &item, err
}
