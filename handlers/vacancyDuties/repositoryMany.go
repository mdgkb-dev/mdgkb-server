package vacancyDuties

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) createMany(items models.VacancyDuties) (err error) {
	_, err = r.db.NewInsert().Model(&items).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []uuid.UUID) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.VacancyDuty)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.VacancyDuties) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("id = EXCLUDED.id").
		Set("name = EXCLUDED.name").
		Set("vacancy_duty_order = EXCLUDED.vacancy_duty_order").
		Model(&items).
		Exec(r.ctx)
	return err
}
