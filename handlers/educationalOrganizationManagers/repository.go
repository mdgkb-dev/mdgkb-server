package educationalOrganizationManagers

import (
	"mdgkb/mdgkb-server/models"

	"github.com/uptrace/bun"
)

func (r *Repository) db() *bun.DB {
	return r.helper.DB.DB
}

func (r *Repository) getAll() (models.EducationalManagers, error) {
	items := make(models.EducationalManagers, 0)
	err := r.db().NewSelect().
		Model(&items).
		Relation("Doctor.Human").
		//Relation("Doctor.FileInfo").
		Order("educational_managers_view.educational_manager_order").
		Scan(r.ctx)
	return items, err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db().NewDelete().
		Model((*models.EducationalManager)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.EducationalManagers) (err error) {
	_, err = r.db().NewInsert().On("conflict (id) do update").
		Set("doctor_id = EXCLUDED.doctor_id").
		Model(&items).
		Exec(r.ctx)
	return err
}
