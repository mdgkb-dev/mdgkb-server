package specializations

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) getAll() (models.Specializations, error) {
	items := make(models.Specializations, 0)
	err := r.db.NewSelect().Model(&items).
		Order("name").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Specialization, error) {
	item := models.Specialization{}
	err := r.db.NewSelect().Model(&item).
		Where("id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.Specialization) (err error) {
	_, err = r.db.NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db.NewDelete().Model(&models.Specialization{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Specialization) (err error) {
	_, err = r.db.NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.Specialization)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Specializations) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("doctor_id = EXCLUDED.doctor_id").
		Set("position = EXCLUDED.position").
		Model(&items).
		Exec(r.ctx)
	return err
}