package appointments

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) getAll() (models.Appointments, error) {
	items := make(models.Appointments, 0)
	err := r.db().NewSelect().Model(&items).
		Relation("Doctor.Human").
		Relation("Doctor.DoctorsDivisions").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) get(id *string) (*models.Appointment, error) {
	item := models.Appointment{}
	err := r.db().NewSelect().Model(&item).
		Relation("Doctor.Human").
		Where("Appointments_view.id = ?", *id).Scan(r.ctx)
	return &item, err
}

func (r *Repository) create(item *models.Appointment) (err error) {
	_, err = r.db().NewInsert().Model(item).Exec(r.ctx)
	return err
}

func (r *Repository) delete(id *string) (err error) {
	_, err = r.db().NewDelete().Model(&models.Appointment{}).Where("id = ?", *id).Exec(r.ctx)
	return err
}

func (r *Repository) update(item *models.Appointment) (err error) {
	_, err = r.db().NewUpdate().Model(item).Where("id = ?", item.ID).Exec(r.ctx)
	return err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.Appointment)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.Appointments) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		//Set("doctor_id = EXCLUDED.doctor_id").
		//Set("position = EXCLUDED.position").
		Model(&items).
		Exec(r.ctx)
	return err
}
