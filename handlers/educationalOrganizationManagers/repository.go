package educationalOrganizationManagers

import (
	"github.com/uptrace/bun"
	"mdgkb/mdgkb-server/models"
)

func (r *Repository) getDB() *bun.DB {
	return r.db
}

func (r *Repository) getAll() (models.EducationalOrganizationManagers, error) {
	items := make(models.EducationalOrganizationManagers, 0)
	err := r.db.NewSelect().Model(&items).
		Relation("Doctor.Human").
		Relation("Doctor.FileInfo").
		Order("educational_organization_managers.manager_order").Scan(r.ctx)
	return items, err
}

func (r *Repository) deleteMany(idPool []string) (err error) {
	_, err = r.db.NewDelete().
		Model((*models.EducationalOrganizationManager)(nil)).
		Where("id IN (?)", bun.In(idPool)).
		Exec(r.ctx)
	return err
}

func (r *Repository) upsertMany(items models.EducationalOrganizationManagers) (err error) {
	_, err = r.db.NewInsert().On("conflict (id) do update").
		Set("doctor_id = EXCLUDED.doctor_id").
		Model(&items).
		Exec(r.ctx)
	return err
}
