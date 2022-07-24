package paidprogramslevels

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
	//_ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.PaidProgramsGroup) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (items models.PaidProgramsGroups, err error) {
	err = r.db().NewSelect().Model(&items).
		Relation("PaidProgramsGroupsGroups.PaidProgramsGroupLevels.YearProgramGroupLevelsGroups.PaidProgramsGroupServices").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(slug string) (*models.PaidProgramsGroup, error) {
	item := models.PaidProgramsGroup{}
	err := r.db().NewSelect().
		Model(&item).
		Relation("PaidProgramsGroupsGroups.PaidProgramsGroupLevels.YearProgramGroupLevelsGroups.PaidProgramsGroupServices").
		Where("year_programs.id = ?", slug).
		Scan(r.ctx)

	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.PaidProgramsGroup{}).Where("id = ?", id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.PaidProgramsGroup) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}
