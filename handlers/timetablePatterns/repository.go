package timetablePatterns

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"

	_ "github.com/go-pg/pg/v10/orm"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) create(item *models.TimetablePattern) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.TimetablePattern) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) getAll() (models.TimetablePatterns, error) {
	items := make(models.TimetablePatterns, 0)
	err := r.db().NewSelect().Model(&items).
		Relation("TimetableDays.Weekday").
		Relation("TimetableDays.BreakPeriods").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id string) (*models.TimetablePattern, error) {
	item := models.TimetablePattern{}
	err := r.db().NewSelect().Model(&item).
		Relation("TimetableDays.Weekday").
		Relation("TimetableDays.BreakPeriods").
		Where("timetable_patterns.id = ?", id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) delete(id string) (err error) {
	_, err = r.db().NewDelete().Model(&models.TimetablePattern{}).Where("id = ?", id).Exec(r.ctx)
	return err
}
