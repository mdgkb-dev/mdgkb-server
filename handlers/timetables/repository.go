package timetables

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"

	_ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) create(item *models.Timetable) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) upsert(item *models.Timetable) (err error) {
	_, err = r.db.NewInsert().Model(item).On("conflict (id) do update").
		Set("description = ?", item.Description).Exec(r.ctx)
	return err
}

func (r *Repository) getAllWeekdays() (items models.Weekdays, err error) {
	err = r.db.NewSelect().Model(&items).Order("number").Scan(r.ctx)
	return items, err
}
