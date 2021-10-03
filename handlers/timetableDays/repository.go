package timetableDays

import (

	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) createMany(items models.TimetableDays) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.Document)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.TimetableDays) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("is_weekend = EXCLUDED.is_weekend").
		Set("timetable_id = EXCLUDED.timetable_id").
		Set("weekday_id = EXCLUDED.weekday_id").
		Set("start_time = EXCLUDED.start_time").
		Set("end_time = EXCLUDED.end_time").
		Set("break_exist = EXCLUDED.break_exist").
		Set("break_start_time = EXCLUDED.break_start_time").
		Set("break_end_time = EXCLUDED.break_end_time").
		Where("timetable_day.id = EXCLUDED.id").
		Exec(r.ctx)
	return err
}
