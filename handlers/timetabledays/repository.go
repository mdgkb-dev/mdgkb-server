package timetabledays

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) createMany(items models.TimetableDays) (err error) {
	_, err = r.db().NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.TimetableDay)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.TimetableDays) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Model(&items).
		Set("is_weekend = EXCLUDED.is_weekend").
		Set("timetable_id = EXCLUDED.timetable_id").
		Set("weekday_id = EXCLUDED.weekday_id").
		Set("start_time = EXCLUDED.start_time").
		Set("end_time = EXCLUDED.end_time").
		Set("breaks_exists = EXCLUDED.breaks_exists").
		Set("around_the_clock = EXCLUDED.around_the_clock").
		Where("timetable_day.id = EXCLUDED.id").
		Exec(r.ctx)
	return err
}
